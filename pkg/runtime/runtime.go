package runtime

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/helton-godoy/shantilly-runtime/pkg/config"
)

// Options represents runtime configuration options
type Options struct {
	Debug   bool
	Verbose bool
}

// Runtime represents the main TUI runtime
type Runtime struct {
	config     *config.Config
	options    Options
	components []Component
	state      *State
}

// Component interface for all UI components
type Component interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
	GetID() string
	SetState(state *State)
}

// State represents the application state
type State struct {
	data       map[string]interface{}
	focus      string
	modalStack []string
	eventQueue []Event
	history    []HistoryEntry
}

// Event represents an application event
type Event struct {
	Type      string
	Source    string
	Data      map[string]interface{}
	Timestamp int64
}

// HistoryEntry represents a state history entry
type HistoryEntry struct {
	Action string
	Data   map[string]interface{}
	Time   int64
}

// New creates a new runtime instance
func New(cfg *config.Config, opts Options) *Runtime {
	return &Runtime{
		config:  cfg,
		options: opts,
		state:   NewState(),
	}
}

// NewState creates a new application state
func NewState() *State {
	return &State{
		data:       make(map[string]interface{}),
		focus:      "",
		modalStack: []string{},
		eventQueue: []Event{},
		history:    []HistoryEntry{},
	}
}

// Init initializes the runtime
func (r *Runtime) Init() tea.Cmd {
	// Initialize components from configuration
	r.components = r.createComponents()

	// Set initial focus
	if len(r.components) > 0 {
		r.state.focus = r.components[0].GetID()
	}

	// Return initial command
	return tea.Batch(
		r.initializeComponents(),
		r.loadInitialState(),
	)
}

// Update handles application updates
func (r *Runtime) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return r.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		return r.handleWindowSizeMsg(msg)
	default:
		return r.handleComponentMsg(msg)
	}
}

// View renders the application
func (r *Runtime) View() string {
	if r.options.Debug {
		return r.debugView()
	}
	return r.renderView()
}

// handleKeyMsg handles keyboard messages
func (r *Runtime) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle global keys
	switch msg.Type {
	case tea.KeyCtrlC:
		return r, tea.Quit
	case tea.KeyTab:
		return r.handleTabKey()
	case tea.KeyEnter:
		return r.handleEnterKey()
	case tea.KeyEsc:
		return r.handleEscKey()
	}

	// Forward to focused component
	if r.state.focus != "" {
		if component := r.getComponent(r.state.focus); component != nil {
			return component.Update(msg)
		}
	}

	return r, nil
}

// handleWindowSizeMsg handles window resize messages
func (r *Runtime) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	// Update all components with new size
	var cmds []tea.Cmd
	for _, component := range r.components {
		_, cmd := component.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return r, tea.Batch(cmds...)
}

// handleComponentMsg handles component-specific messages
func (r *Runtime) handleComponentMsg(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Forward to all components
	var cmds []tea.Cmd
	for _, component := range r.components {
		_, cmd := component.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return r, tea.Batch(cmds...)
}

// handleTabKey handles tab navigation
func (r *Runtime) handleTabKey() (tea.Model, tea.Cmd) {
	// Find next focusable component
	currentIndex := -1
	for i, component := range r.components {
		if component.GetID() == r.state.focus {
			currentIndex = i
			break
		}
	}

	// Move to next component
	nextIndex := (currentIndex + 1) % len(r.components)
	r.state.focus = r.components[nextIndex].GetID()

	return r, nil
}

// handleEnterKey handles enter key
func (r *Runtime) handleEnterKey() (tea.Model, tea.Cmd) {
	// Trigger click event on focused component
	if r.state.focus != "" {
		event := Event{
			Type:   "click",
			Source: r.state.focus,
			Data:   map[string]interface{}{},
		}

		return r.handleEvent(event)
	}

	return r, nil
}

// handleEscKey handles escape key
func (r *Runtime) handleEscKey() (tea.Model, tea.Cmd) {
	// Close modal if open
	if len(r.state.modalStack) > 0 {
		r.state.modalStack = r.state.modalStack[:len(r.state.modalStack)-1]
		return r, nil
	}

	return r, nil
}

// handleEvent processes application events
func (r *Runtime) handleEvent(event Event) (tea.Model, tea.Cmd) {
	// Add to history
	r.addHistory("event", map[string]interface{}{
		"type":   event.Type,
		"source": event.Source,
		"data":   event.Data,
	})

	// Find component configuration
	component, err := r.config.GetComponentByID(event.Source)
	if err != nil {
		if r.options.Verbose {
			fmt.Printf("Warning: Component %s not found\n", event.Source)
		}
		return r, nil
	}

	// Process component events
	if component.Events != nil {
		if eventAction, exists := component.Events[event.Type]; exists {
			return r.executeAction(eventAction)
		}
	}

	return r, nil
}

// executeAction executes an action from configuration
func (r *Runtime) executeAction(action interface{}) (tea.Model, tea.Cmd) {
	// Convert action to map
	actionMap, ok := action.(map[string]interface{})
	if !ok {
		return r, nil
	}

	actionType, ok := actionMap["type"].(string)
	if !ok {
		return r, nil
	}

	switch actionType {
	case "run":
		return r.executeRunAction(actionMap)
	case "navigate":
		return r.executeNavigateAction(actionMap)
	case "show-modal":
		return r.executeShowModalAction(actionMap)
	case "hide-modal":
		return r.executeHideModalAction(actionMap)
	case "update":
		return r.executeUpdateAction(actionMap)
	case "conditional":
		return r.executeConditionalAction(actionMap)
	case "exit":
		return r.executeExitAction(actionMap)
	default:
		if r.options.Verbose {
			fmt.Printf("Warning: Unknown action type: %s\n", actionType)
		}
	}

	return r, nil
}

// executeRunAction executes a run command action
func (r *Runtime) executeRunAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	command, ok := actionMap["command"].(string)
	if !ok {
		return r, nil
	}

	// Execute command (in real implementation, this would be more sophisticated)
	if r.options.Verbose {
		fmt.Printf("Executing command: %s\n", command)
	}

	// For now, just log the action
	r.addHistory("run", map[string]interface{}{
		"command": command,
	})

	return r, nil
}

// executeNavigateAction handles navigation actions
func (r *Runtime) executeNavigateAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	target, ok := actionMap["target"].(string)
	if !ok {
		return r, nil
	}

	r.addHistory("navigate", map[string]interface{}{
		"target": target,
	})

	return r, nil
}

// executeShowModalAction shows a modal
func (r *Runtime) executeShowModalAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	modalID, ok := actionMap["modal"].(string)
	if !ok {
		return r, nil
	}

	r.state.modalStack = append(r.state.modalStack, modalID)

	return r, nil
}

// executeHideModalAction hides a modal
func (r *Runtime) executeHideModalAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	if len(r.state.modalStack) > 0 {
		r.state.modalStack = r.state.modalStack[:len(r.state.modalStack)-1]
	}

	return r, nil
}

// executeUpdateAction updates component data
func (r *Runtime) executeUpdateAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	target, ok := actionMap["target"].(string)
	if !ok {
		return r, nil
	}

	value, ok := actionMap["value"]
	if !ok {
		return r, nil
	}

	r.state.data[target] = value

	return r, nil
}

// executeConditionalAction handles conditional actions
func (r *Runtime) executeConditionalAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	conditions, ok := actionMap["conditions"].([]interface{})
	if !ok {
		return r, nil
	}

	// Evaluate conditions
	for _, condition := range conditions {
		conditionMap, ok := condition.(map[string]interface{})
		if !ok {
			continue
		}

		when, ok := conditionMap["when"].(string)
		if !ok {
			continue
		}

		// Simple condition evaluation (would be more sophisticated in real implementation)
		if r.evaluateCondition(when) {
			if thenAction, exists := conditionMap["then"]; exists {
				return r.executeAction(thenAction)
			}
		}
	}

	return r, nil
}

// executeExitAction exits the application
func (r *Runtime) executeExitAction(actionMap map[string]interface{}) (tea.Model, tea.Cmd) {
	code, ok := actionMap["code"].(int)
	if !ok {
		code = 0
	}

	r.addHistory("exit", map[string]interface{}{
		"code": code,
	})

	return r, tea.Quit
}

// evaluateCondition evaluates a condition string
func (r *Runtime) evaluateCondition(condition string) bool {
	// Simple condition evaluation (would be more sophisticated)
	// For now, just return true for demonstration
	return true
}

// createComponents creates components from configuration
func (r *Runtime) createComponents() []Component {
	var components []Component

	for _, configComponent := range r.config.Layout.Components {
		component := r.createComponent(configComponent)
		if component != nil {
			components = append(components, component)
		}
	}

	return components
}

// createComponent creates a single component from configuration
func (r *Runtime) createComponent(configComponent config.Component) Component {
	switch configComponent.Type {
	case "text":
		return NewTextComponent(configComponent, r.state)
	case "button":
		return NewButtonComponent(configComponent, r.state)
	case "input":
		return NewInputComponent(configComponent, r.state)
	case "select":
		return NewSelectComponent(configComponent, r.state)
	case "progress":
		return NewProgressComponent(configComponent, r.state)
	case "spinner":
		return NewSpinnerComponent(configComponent, r.state)
	case "table":
		return NewTableComponent(configComponent, r.state)
	case "modal":
		return NewModalComponent(configComponent, r.state)
	default:
		if r.options.Verbose {
			fmt.Printf("Warning: Unknown component type: %s\n", configComponent.Type)
		}
		return nil
	}
}

// getComponent finds a component by ID
func (r *Runtime) getComponent(id string) Component {
	for _, component := range r.components {
		if component.GetID() == id {
			return component
		}
	}
	return nil
}

// initializeComponents initializes all components
func (r *Runtime) initializeComponents() tea.Cmd {
	var cmds []tea.Cmd

	for _, component := range r.components {
		cmd := component.Init()
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return tea.Batch(cmds...)
}

// loadInitialState loads the initial application state
func (r *Runtime) loadInitialState() tea.Cmd {
	// Load initial state from configuration or defaults
	return nil
}

// renderView renders the main view
func (r *Runtime) renderView() string {
	// Render components based on layout
	var view string

	for _, component := range r.components {
		view += component.View() + "\n"
	}

	return view
}

// debugView renders a debug view
func (r *Runtime) debugView() string {
	view := "=== DEBUG VIEW ===\n"
	view += fmt.Sprintf("Focus: %s\n", r.state.focus)
	view += fmt.Sprintf("Components: %d\n", len(r.components))
	view += fmt.Sprintf("Modal Stack: %v\n", r.state.modalStack)
	view += fmt.Sprintf("Data: %+v\n", r.state.data)
	view += fmt.Sprintf("History: %d entries\n", len(r.state.history))
	view += "\n=== COMPONENTS ===\n"

	for _, component := range r.components {
		view += fmt.Sprintf("ID: %s, Type: %T\n", component.GetID(), component)
	}

	view += "\n=== NORMAL VIEW ===\n"
	view += r.renderView()

	return view
}

// addHistory adds an entry to the history
func (r *Runtime) addHistory(action string, data map[string]interface{}) {
	entry := HistoryEntry{
		Action: action,
		Data:   data,
		Time:   0, // Would use actual timestamp
	}

	r.state.history = append(r.state.history, entry)

	// Limit history size
	if len(r.state.history) > 100 {
		r.state.history = r.state.history[1:]
	}
}

// Run starts the TUI application
func (r *Runtime) Run() error {
	program := tea.NewProgram(r, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		return fmt.Errorf("failed to start TUI: %w", err)
	}

	return nil
}

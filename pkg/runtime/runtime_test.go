package runtime

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/helton-godoy/shantilly-runtime/pkg/config"
)

func TestNewRuntime(t *testing.T) {
	cfg := config.DefaultConfig()
	opts := Options{
		Debug:   true,
		Verbose: true,
	}

	rt := New(cfg, opts)

	if rt.config == nil {
		t.Error("Runtime should have config")
	}

	if rt.options.Debug != true {
		t.Error("Runtime should have debug enabled")
	}

	if rt.options.Verbose != true {
		t.Error("Runtime should have verbose enabled")
	}

	if rt.state == nil {
		t.Error("Runtime should have state")
	}
}

func TestNewState(t *testing.T) {
	state := NewState()

	if state.data == nil {
		t.Error("State should have data map")
	}

	if state.modalStack == nil {
		t.Error("State should have modal stack")
	}

	if state.eventQueue == nil {
		t.Error("State should have event queue")
	}

	if state.history == nil {
		t.Error("State should have history")
	}
}

func TestRuntimeInit(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})

	cmd := rt.Init()

	// Init may return nil command, that's okay
	if cmd == nil {
		t.Log("Init returned nil command (acceptable)")
	}

	if len(rt.components) == 0 {
		t.Error("Runtime should have components after init")
	}

	// Focus may be empty if no component has ID
	if len(rt.components) > 0 && rt.components[0].GetID() != "" && rt.state.focus == "" {
		t.Error("Runtime should have focus set after init when components have IDs")
	}
}

func TestRuntimeUpdateKeyMsg(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})
	rt.Init()

	// Test Tab key
	model, _ := rt.Update(tea.KeyMsg{Type: tea.KeyTab})

	if model == nil {
		t.Error("Update should return a model")
	}

	// Test Ctrl+C key
	model, cmd := rt.Update(tea.KeyMsg{Type: tea.KeyCtrlC})

	if cmd == nil {
		t.Error("Ctrl+C should return quit command")
	}

	// Use cmd to avoid unused variable error
	_ = cmd
}

func TestRuntimeUpdateWindowSizeMsg(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})
	rt.Init()

	msg := tea.WindowSizeMsg{
		Width:  80,
		Height: 24,
	}

	model, cmd := rt.Update(msg)

	if model == nil {
		t.Error("Update should return a model")
	}

	// Use cmd to avoid unused variable error
	_ = cmd
}

func TestGetComponent(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})
	rt.Init()

	// Test existing component
	if len(rt.components) > 0 {
		firstComponent := rt.components[0]
		found := rt.getComponent(firstComponent.GetID())

		if found == nil {
			t.Error("Should find existing component")
		}

		if found.GetID() != firstComponent.GetID() {
			t.Error("Found component should have correct ID")
		}
	}

	// Test non-existing component
	found := rt.getComponent("non-existent")
	if found != nil {
		t.Error("Should not find non-existent component")
	}
}

func TestAddHistory(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})

	initialHistorySize := len(rt.state.history)

	rt.addHistory("test", map[string]interface{}{
		"key": "value",
	})

	if len(rt.state.history) != initialHistorySize+1 {
		t.Error("History should have one more entry")
	}

	lastEntry := rt.state.history[len(rt.state.history)-1]
	if lastEntry.Action != "test" {
		t.Error("Last entry should have correct action")
	}
}

func TestCreateComponent(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})

	// Test text component
	textConfig := config.Component{
		Type:    "text",
		Content: "Hello World",
		Style:   "bold",
	}

	component := rt.createComponent(textConfig)

	if component == nil {
		t.Error("Should create text component")
	}

	if component.GetID() != "" {
		t.Error("Text component should not have ID by default")
	}

	// Test button component
	buttonConfig := config.Component{
		ID:      "test-btn",
		Type:    "button",
		Content: "Click Me",
	}

	component = rt.createComponent(buttonConfig)

	if component == nil {
		t.Error("Should create button component")
	}

	if component.GetID() != "test-btn" {
		t.Error("Button component should have correct ID")
	}

	// Test unknown component type
	unknownConfig := config.Component{
		Type: "unknown",
	}

	component = rt.createComponent(unknownConfig)

	if component != nil {
		t.Error("Should not create unknown component type")
	}
}

func TestExecuteAction(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})

	// Test run action
	action := map[string]interface{}{
		"type":    "run",
		"command": "echo test",
	}

	model, cmd := rt.executeAction(action)

	if model == nil {
		t.Error("ExecuteAction should return a model")
	}

	if cmd != nil {
		t.Error("Run action should not return command")
	}

	// Test exit action
	action = map[string]interface{}{
		"type": "exit",
		"code": 0,
	}

	model, cmd = rt.executeAction(action)

	if cmd == nil {
		t.Error("Exit action should return quit command")
	}

	// Test unknown action
	action = map[string]interface{}{
		"type": "unknown",
	}

	model, cmd = rt.executeAction(action)

	if model == nil {
		t.Error("ExecuteAction should return a model for unknown action")
	}

	if cmd != nil {
		t.Error("Unknown action should not return command")
	}
}

func TestEvaluateCondition(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})

	// For now, evaluateCondition always returns true
	result := rt.evaluateCondition("any condition")

	if !result {
		t.Error("EvaluateCondition should return true for now")
	}
}

func TestRenderView(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{})
	rt.Init()

	view := rt.renderView()

	if view == "" {
		t.Error("RenderView should return non-empty string")
	}
}

func TestDebugView(t *testing.T) {
	cfg := config.DefaultConfig()
	rt := New(cfg, Options{
		Debug:   true,
		Verbose: true,
	})
	rt.Init()

	view := rt.debugView()

	if view == "" {
		t.Error("DebugView should return non-empty string")
	}

	if !contains(view, "DEBUG VIEW") {
		t.Error("DebugView should contain debug header")
	}
}

func TestTextComponent(t *testing.T) {
	cfg := config.Component{
		Type:    "text",
		Content: "Hello World",
		Style:   "bold",
	}

	state := NewState()
	component := NewTextComponent(cfg, state)

	if component == nil {
		t.Error("Should create text component")
	}

	if component.GetID() != "" {
		t.Error("Text component should not have ID by default")
	}

	view := component.View()

	if view == "" {
		t.Error("Text component view should not be empty")
	}

	if !contains(view, "Hello World") {
		t.Error("Text component view should contain content")
	}
}

func TestButtonComponent(t *testing.T) {
	cfg := config.Component{
		ID:      "test-btn",
		Type:    "button",
		Content: "Click Me",
	}

	state := NewState()
	component := NewButtonComponent(cfg, state)

	if component == nil {
		t.Error("Should create button component")
	}

	if component.GetID() != "test-btn" {
		t.Error("Button component should have correct ID")
	}

	view := component.View()

	if view == "" {
		t.Error("Button component view should not be empty")
	}

	if !contains(view, "Click Me") {
		t.Error("Button component view should contain content")
	}
}

func TestInputComponent(t *testing.T) {
	cfg := config.Component{
		ID:   "test-input",
		Type: "input",
		Props: map[string]interface{}{
			"placeholder": "Enter text...",
		},
	}

	state := NewState()
	component := NewInputComponent(cfg, state)

	if component == nil {
		t.Error("Should create input component")
	}

	if component.GetID() != "test-input" {
		t.Error("Input component should have correct ID")
	}

	view := component.View()

	if view == "" {
		t.Error("Input component view should not be empty")
	}
}

func TestSelectComponent(t *testing.T) {
	cfg := config.Component{
		ID:   "test-select",
		Type: "select",
		Props: map[string]interface{}{
			"options": []interface{}{"Option 1", "Option 2", "Option 3"},
		},
	}

	state := NewState()
	component := NewSelectComponent(cfg, state)

	if component == nil {
		t.Error("Should create select component")
	}

	if component.GetID() != "test-select" {
		t.Error("Select component should have correct ID")
	}

	view := component.View()

	if view == "" {
		t.Error("Select component view should not be empty")
	}
}

func TestProgressComponent(t *testing.T) {
	cfg := config.Component{
		ID:   "test-progress",
		Type: "progress",
		Props: map[string]interface{}{
			"max": 100,
		},
	}

	state := NewState()
	component := NewProgressComponent(cfg, state)

	if component == nil {
		t.Error("Should create progress component")
	}

	if component.GetID() != "test-progress" {
		t.Error("Progress component should have correct ID")
	}

	view := component.View()

	if view == "" {
		t.Error("Progress component view should not be empty")
	}
}

func TestSpinnerComponent(t *testing.T) {
	cfg := config.Component{
		ID:   "test-spinner",
		Type: "spinner",
	}

	state := NewState()
	component := NewSpinnerComponent(cfg, state)

	if component == nil {
		t.Error("Should create spinner component")
	}

	if component.GetID() != "test-spinner" {
		t.Error("Spinner component should have correct ID")
	}

	view := component.View()

	if view == "" {
		t.Error("Spinner component view should not be empty")
	}
}

func TestTableComponent(t *testing.T) {
	cfg := config.Component{
		ID:   "test-table",
		Type: "table",
	}

	state := NewState()
	component := NewTableComponent(cfg, state)

	if component == nil {
		t.Error("Should create table component")
	}

	if component.GetID() != "test-table" {
		t.Error("Table component should have correct ID")
	}

	view := component.View()

	if view == "" {
		t.Error("Table component view should not be empty")
	}
}

func TestModalComponent(t *testing.T) {
	cfg := config.Component{
		ID:   "test-modal",
		Type: "modal",
	}

	state := NewState()
	component := NewModalComponent(cfg, state)

	if component == nil {
		t.Error("Should create modal component")
	}

	if component.GetID() != "test-modal" {
		t.Error("Modal component should have correct ID")
	}

	// Test hidden modal
	view := component.View()
	if view != "" {
		t.Error("Hidden modal should return empty view")
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			indexOf(s, substr) >= 0))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

package runtime

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/helton-godoy/shantilly-runtime/pkg/config"
)

// TextComponent represents a text display component
type TextComponent struct {
	id     string
	config config.Component
	state  *State
	style  lipgloss.Style
}

// NewTextComponent creates a new text component
func NewTextComponent(cfg config.Component, state *State) *TextComponent {
	return &TextComponent{
		id:     cfg.ID,
		config: cfg,
		state:  state,
		style:  parseStyle(cfg.Style),
	}
}

// Init initializes the text component
func (t *TextComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the text component
func (t *TextComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

// View renders the text component
func (t *TextComponent) View() string {
	content := t.config.Content
	if content == "" {
		content = ""
	}

	return t.style.Render(content)
}

// GetID returns the component ID
func (t *TextComponent) GetID() string {
	return t.id
}

// SetState sets the component state
func (t *TextComponent) SetState(state *State) {
	t.state = state
}

// ButtonComponent represents a button component
type ButtonComponent struct {
	id      string
	config  config.Component
	state   *State
	style   lipgloss.Style
	focused bool
}

// NewButtonComponent creates a new button component
func NewButtonComponent(cfg config.Component, state *State) *ButtonComponent {
	return &ButtonComponent{
		id:      cfg.ID,
		config:  cfg,
		state:   state,
		style:   parseStyle(cfg.Style),
		focused: false,
	}
}

// Init initializes the button component
func (b *ButtonComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the button component
func (b *ButtonComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter && b.focused {
			// Trigger click event
			return b, b.handleClick()
		}
	}

	return b, nil
}

// View renders the button component
func (b *ButtonComponent) View() string {
	content := b.config.Content
	if content == "" {
		content = "Button"
	}

	style := b.style
	if b.focused {
		style = style.Background(lipgloss.Color("240"))
	}

	return style.Render(content)
}

// GetID returns the component ID
func (b *ButtonComponent) GetID() string {
	return b.id
}

// SetState sets the component state
func (b *ButtonComponent) SetState(state *State) {
	b.state = state
}

// SetFocus sets the focus state
func (b *ButtonComponent) SetFocus(focused bool) {
	b.focused = focused
}

// handleClick handles button click events
func (b *ButtonComponent) handleClick() tea.Cmd {
	return func() tea.Msg {
		return Event{
			Type:   "click",
			Source: b.id,
			Data:   map[string]interface{}{},
		}
	}
}

// InputComponent represents an input field component
type InputComponent struct {
	id      string
	config  config.Component
	state   *State
	style   lipgloss.Style
	value   string
	focused bool
	cursor  int
}

// NewInputComponent creates a new input component
func NewInputComponent(cfg config.Component, state *State) *InputComponent {
	return &InputComponent{
		id:      cfg.ID,
		config:  cfg,
		state:   state,
		style:   parseStyle(cfg.Style),
		value:   "",
		focused: false,
		cursor:  0,
	}
}

// Init initializes the input component
func (i *InputComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the input component
func (i *InputComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !i.focused {
			return i, nil
		}

		switch msg.Type {
		case tea.KeyRunes:
			i.value += string(msg.Runes)
			i.cursor = len(i.value)
		case tea.KeyBackspace:
			if i.cursor > 0 {
				i.value = i.value[:i.cursor-1] + i.value[i.cursor:]
				i.cursor--
			}
		case tea.KeyDelete:
			if i.cursor < len(i.value) {
				i.value = i.value[:i.cursor] + i.value[i.cursor+1:]
			}
		case tea.KeyLeft:
			if i.cursor > 0 {
				i.cursor--
			}
		case tea.KeyRight:
			if i.cursor < len(i.value) {
				i.cursor++
			}
		case tea.KeyHome:
			i.cursor = 0
		case tea.KeyEnd:
			i.cursor = len(i.value)
		}
	}

	return i, nil
}

// View renders the input component
func (i *InputComponent) View() string {
	placeholder := ""
	if ph, ok := i.config.Props["placeholder"].(string); ok {
		placeholder = ph
	}
	if placeholder == "" {
		placeholder = "Enter text..."
	}

	displayValue := i.value
	if displayValue == "" {
		displayValue = placeholder
	}

	style := i.style
	if i.focused {
		style = style.Background(lipgloss.Color("240"))
	}

	return style.Render(displayValue)
}

// GetID returns the component ID
func (i *InputComponent) GetID() string {
	return i.id
}

// SetState sets the component state
func (i *InputComponent) SetState(state *State) {
	i.state = state
}

// SetFocus sets the focus state
func (i *InputComponent) SetFocus(focused bool) {
	i.focused = focused
}

// GetValue returns the current input value
func (i *InputComponent) GetValue() string {
	return i.value
}

// SetValue sets the input value
func (i *InputComponent) SetValue(value string) {
	i.value = value
	i.cursor = len(value)
}

// SelectComponent represents a select dropdown component
type SelectComponent struct {
	id       string
	config   config.Component
	state    *State
	style    lipgloss.Style
	options  []string
	selected int
	focused  bool
}

// NewSelectComponent creates a new select component
func NewSelectComponent(cfg config.Component, state *State) *SelectComponent {
	options := []string{}
	if opts, ok := cfg.Props["options"].([]interface{}); ok {
		for _, opt := range opts {
			if str, ok := opt.(string); ok {
				options = append(options, str)
			}
		}
	}

	return &SelectComponent{
		id:       cfg.ID,
		config:   cfg,
		state:    state,
		style:    parseStyle(cfg.Style),
		options:  options,
		selected: 0,
		focused:  false,
	}
}

// Init initializes the select component
func (s *SelectComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the select component
func (s *SelectComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !s.focused {
			return s, nil
		}

		switch msg.Type {
		case tea.KeyUp:
			if s.selected > 0 {
				s.selected--
			}
		case tea.KeyDown:
			if s.selected < len(s.options)-1 {
				s.selected++
			}
		case tea.KeyEnter:
			return s, s.handleChange()
		}
	}

	return s, nil
}

// View renders the select component
func (s *SelectComponent) View() string {
	if len(s.options) == 0 {
		displayValue := "No options available"
		style := s.style
		if s.focused {
			style = style.Background(lipgloss.Color("240"))
		}
		return style.Render(displayValue)
	}

	displayValue := s.options[s.selected]
	style := s.style
	if s.focused {
		style = style.Background(lipgloss.Color("240"))
	}

	return style.Render(displayValue)
}

// GetID returns the component ID
func (s *SelectComponent) GetID() string {
	return s.id
}

// SetState sets the component state
func (s *SelectComponent) SetState(state *State) {
	s.state = state
}

// SetFocus sets the focus state
func (s *SelectComponent) SetFocus(focused bool) {
	s.focused = focused
}

// GetValue returns the selected value
func (s *SelectComponent) GetValue() string {
	if s.selected < len(s.options) {
		return s.options[s.selected]
	}
	return ""
}

// handleChange handles selection change events
func (s *SelectComponent) handleChange() tea.Cmd {
	return func() tea.Msg {
		return Event{
			Type:   "change",
			Source: s.id,
			Data: map[string]interface{}{
				"value": s.GetValue(),
			},
		}
	}
}

// ProgressComponent represents a progress bar component
type ProgressComponent struct {
	id     string
	config config.Component
	state  *State
	style  lipgloss.Style
	value  float64
	max    float64
}

// NewProgressComponent creates a new progress component
func NewProgressComponent(cfg config.Component, state *State) *ProgressComponent {
	max := 100.0
	if maxVal, ok := cfg.Props["max"].(int); ok {
		max = float64(maxVal)
	}

	return &ProgressComponent{
		id:     cfg.ID,
		config: cfg,
		state:  state,
		style:  parseStyle(cfg.Style),
		value:  0,
		max:    max,
	}
}

// Init initializes the progress component
func (p *ProgressComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the progress component
func (p *ProgressComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

// View renders the progress component
func (p *ProgressComponent) View() string {
	percentage := p.value / p.max
	if percentage > 1 {
		percentage = 1
	}
	if percentage < 0 {
		percentage = 0
	}

	width := 20
	filled := int(percentage * float64(width))

	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return p.style.Render(fmt.Sprintf("[%s] %.0f%%", bar, percentage*100))
}

// GetID returns the component ID
func (p *ProgressComponent) GetID() string {
	return p.id
}

// SetState sets the component state
func (p *ProgressComponent) SetState(state *State) {
	p.state = state
}

// SetValue sets the progress value
func (p *ProgressComponent) SetValue(value float64) {
	p.value = value
}

// SpinnerComponent represents a spinner/loading component
type SpinnerComponent struct {
	id     string
	config config.Component
	state  *State
	style  lipgloss.Style
	frame  int
	frames []string
}

// NewSpinnerComponent creates a new spinner component
func NewSpinnerComponent(cfg config.Component, state *State) *SpinnerComponent {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	return &SpinnerComponent{
		id:     cfg.ID,
		config: cfg,
		state:  state,
		style:  parseStyle(cfg.Style),
		frame:  0,
		frames: frames,
	}
}

// Init initializes the spinner component
func (s *SpinnerComponent) Init() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update handles updates for the spinner component
func (s *SpinnerComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TickMsg:
		s.frame = (s.frame + 1) % len(s.frames)
		return s, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
			return TickMsg(t)
		})
	}

	return s, nil
}

// View renders the spinner component
func (s *SpinnerComponent) View() string {
	return s.style.Render(s.frames[s.frame])
}

// GetID returns the component ID
func (s *SpinnerComponent) GetID() string {
	return s.id
}

// SetState sets the component state
func (s *SpinnerComponent) SetState(state *State) {
	s.state = state
}

// TableComponent represents a table component
type TableComponent struct {
	id      string
	config  config.Component
	state   *State
	style   lipgloss.Style
	data    [][]string
	headers []string
}

// NewTableComponent creates a new table component
func NewTableComponent(cfg config.Component, state *State) *TableComponent {
	return &TableComponent{
		id:      cfg.ID,
		config:  cfg,
		state:   state,
		style:   parseStyle(cfg.Style),
		data:    [][]string{},
		headers: []string{},
	}
}

// Init initializes the table component
func (t *TableComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the table component
func (t *TableComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

// View renders the table component
func (t *TableComponent) View() string {
	if len(t.data) == 0 && len(t.headers) == 0 {
		return t.style.Render("No data available")
	}

	result := ""

	// Render headers
	if len(t.headers) > 0 {
		result += t.style.Render(fmt.Sprintf("%-20s", t.headers[0]))
		for _, header := range t.headers[1:] {
			result += t.style.Render(fmt.Sprintf(" | %-20s", header))
		}
		result += "\n"
	}

	// Render data rows
	for _, row := range t.data {
		if len(row) > 0 {
			result += t.style.Render(fmt.Sprintf("%-20s", row[0]))
			for _, cell := range row[1:] {
				result += t.style.Render(fmt.Sprintf(" | %-20s", cell))
			}
			result += "\n"
		}
	}

	return result
}

// GetID returns the component ID
func (t *TableComponent) GetID() string {
	return t.id
}

// SetState sets the component state
func (t *TableComponent) SetState(state *State) {
	t.state = state
}

// SetData sets the table data
func (t *TableComponent) SetData(headers []string, data [][]string) {
	t.headers = headers
	t.data = data
}

// ModalComponent represents a modal dialog component
type ModalComponent struct {
	id      string
	config  config.Component
	state   *State
	style   lipgloss.Style
	visible bool
	title   string
	content string
}

// NewModalComponent creates a new modal component
func NewModalComponent(cfg config.Component, state *State) *ModalComponent {
	return &ModalComponent{
		id:      cfg.ID,
		config:  cfg,
		state:   state,
		style:   parseStyle(cfg.Style),
		visible: false,
		title:   "Modal",
		content: "",
	}
}

// Init initializes the modal component
func (m *ModalComponent) Init() tea.Cmd {
	return nil
}

// Update handles updates for the modal component
func (m *ModalComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.visible && msg.Type == tea.KeyEsc {
			m.visible = false
		}
	}

	return m, nil
}

// View renders the modal component
func (m *ModalComponent) View() string {
	if !m.visible {
		return ""
	}

	modalStyle := m.style.
		Width(40).
		Height(10).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder())

	content := fmt.Sprintf("%s\n\n%s", m.title, m.content)

	return modalStyle.Render(content)
}

// GetID returns the component ID
func (m *ModalComponent) GetID() string {
	return m.id
}

// SetState sets the component state
func (m *ModalComponent) SetState(state *State) {
	m.state = state
}

// Show shows the modal
func (m *ModalComponent) Show(title, content string) {
	m.visible = true
	m.title = title
	m.content = content
}

// Hide hides the modal
func (m *ModalComponent) Hide() {
	m.visible = false
}

// IsVisible returns whether the modal is visible
func (m *ModalComponent) IsVisible() bool {
	return m.visible
}

// parseStyle parses a style string into a lipgloss Style
func parseStyle(styleStr string) lipgloss.Style {
	style := lipgloss.NewStyle()

	if styleStr == "" {
		return style
	}

	// Parse simple style options
	if styleStr == "bold" {
		style = style.Bold(true)
	} else if styleStr == "italic" {
		style = style.Italic(true)
	} else if styleStr == "underline" {
		style = style.Underline(true)
	}

	return style
}

// TickMsg represents a timer tick message
type TickMsg time.Time

package config

import (
	"fmt"
	"strings"
)

// ComponentType represents the type of a UI component
type ComponentType string

const (
	ComponentTypeText     ComponentType = "text"
	ComponentTypeButton   ComponentType = "button"
	ComponentTypeInput    ComponentType = "input"
	ComponentTypeSelect   ComponentType = "select"
	ComponentTypeProgress ComponentType = "progress"
	ComponentTypeSpinner  ComponentType = "spinner"
	ComponentTypeTable    ComponentType = "table"
	ComponentTypeModal    ComponentType = "modal"
)

// LayoutType represents the type of layout
type LayoutType string

const (
	LayoutTypeVertical   LayoutType = "vertical"
	LayoutTypeHorizontal LayoutType = "horizontal"
	LayoutTypeGrid       LayoutType = "grid"
	LayoutTypeModal      LayoutType = "modal"
)

// EventType represents the type of event a component can handle
type EventType string

const (
	EventTypeClick    EventType = "click"
	EventTypeChange   EventType = "change"
	EventTypeFocus    EventType = "focus"
	EventTypeBlur     EventType = "blur"
	EventTypeKeyPress EventType = "keypress"
	EventTypeSubmit   EventType = "submit"
	EventTypeLoad     EventType = "load"
	EventTypeUnload   EventType = "unload"
)

// ActionType represents the type of action to perform
type ActionType string

const (
	ActionTypeRun         ActionType = "run"
	ActionTypeNavigate    ActionType = "navigate"
	ActionTypeShowModal   ActionType = "show-modal"
	ActionTypeHideModal   ActionType = "hide-modal"
	ActionTypeUpdate      ActionType = "update"
	ActionTypeConditional ActionType = "conditional"
	ActionTypeExit        ActionType = "exit"
)

// Style represents styling options for components
type Style struct {
	Color      string `yaml:"color,omitempty" json:"color,omitempty"`
	Background string `yaml:"background,omitempty" json:"background,omitempty"`
	Bold       bool   `yaml:"bold,omitempty" json:"bold,omitempty"`
	Italic     bool   `yaml:"italic,omitempty" json:"italic,omitempty"`
	Underline  bool   `yaml:"underline,omitempty" json:"underline,omitempty"`
	Width      int    `yaml:"width,omitempty" json:"width,omitempty"`
	Height     int    `yaml:"height,omitempty" json:"height,omitempty"`
	Align      string `yaml:"align,omitempty" json:"align,omitempty"`
}

// Props represents component properties
type Props struct {
	Placeholder string                 `yaml:"placeholder,omitempty" json:"placeholder,omitempty"`
	Required    bool                   `yaml:"required,omitempty" json:"required,omitempty"`
	Disabled    bool                   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
	Options     []string               `yaml:"options,omitempty" json:"options,omitempty"`
	Value       interface{}            `yaml:"value,omitempty" json:"value,omitempty"`
	Min         int                    `yaml:"min,omitempty" json:"min,omitempty"`
	Max         int                    `yaml:"max,omitempty" json:"max,omitempty"`
	Step        int                    `yaml:"step,omitempty" json:"step,omitempty"`
	Custom      map[string]interface{} `yaml:"custom,omitempty" json:"custom,omitempty"`
}

// EventAction represents an action to be performed when an event occurs
type EventAction struct {
	Type         ActionType             `yaml:"type" json:"type"`
	Command      string                 `yaml:"command,omitempty" json:"command,omitempty"`
	Target       string                 `yaml:"target,omitempty" json:"target,omitempty"`
	Conditions   []EventCondition       `yaml:"conditions,omitempty" json:"conditions,omitempty"`
	Parameters   map[string]interface{} `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	UpdateTarget string                 `yaml:"update_target,omitempty" json:"update_target,omitempty"`
	Code         int                    `yaml:"code,omitempty" json:"code,omitempty"`
}

// EventCondition represents a condition for conditional actions
type EventCondition struct {
	When string       `yaml:"when" json:"when"`
	Then *EventAction `yaml:"then,omitempty" json:"then,omitempty"`
	Else *EventAction `yaml:"else,omitempty" json:"else,omitempty"`
}

// ValidateComponentType validates if a component type is supported
func ValidateComponentType(componentType string) error {
	validTypes := []ComponentType{
		ComponentTypeText,
		ComponentTypeButton,
		ComponentTypeInput,
		ComponentTypeSelect,
		ComponentTypeProgress,
		ComponentTypeSpinner,
		ComponentTypeTable,
		ComponentTypeModal,
	}

	for _, t := range validTypes {
		if string(t) == componentType {
			return nil
		}
	}

	return fmt.Errorf("unsupported component type: %s", componentType)
}

// ValidateLayoutType validates if a layout type is supported
func ValidateLayoutType(layoutType string) error {
	validTypes := []LayoutType{
		LayoutTypeVertical,
		LayoutTypeHorizontal,
		LayoutTypeGrid,
		LayoutTypeModal,
	}

	for _, t := range validTypes {
		if string(t) == layoutType {
			return nil
		}
	}

	return fmt.Errorf("unsupported layout type: %s", layoutType)
}

// ValidateEventType validates if an event type is supported
func ValidateEventType(eventType string) error {
	validTypes := []EventType{
		EventTypeClick,
		EventTypeChange,
		EventTypeFocus,
		EventTypeBlur,
		EventTypeKeyPress,
		EventTypeSubmit,
		EventTypeLoad,
		EventTypeUnload,
	}

	for _, t := range validTypes {
		if string(t) == eventType {
			return nil
		}
	}

	return fmt.Errorf("unsupported event type: %s", eventType)
}

// ValidateActionType validates if an action type is supported
func ValidateActionType(actionType string) error {
	validTypes := []ActionType{
		ActionTypeRun,
		ActionTypeNavigate,
		ActionTypeShowModal,
		ActionTypeHideModal,
		ActionTypeUpdate,
		ActionTypeConditional,
		ActionTypeExit,
	}

	for _, t := range validTypes {
		if string(t) == actionType {
			return nil
		}
	}

	return fmt.Errorf("unsupported action type: %s", actionType)
}

// ParseStyleString parses a style string into Style struct
func ParseStyleString(styleStr string) Style {
	style := Style{}

	if styleStr == "" {
		return style
	}

	parts := strings.Split(styleStr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)

		switch part {
		case "bold":
			style.Bold = true
		case "italic":
			style.Italic = true
		case "underline":
			style.Underline = true
		default:
			// Handle color and other properties
			if strings.HasPrefix(part, "color=") {
				style.Color = strings.TrimPrefix(part, "color=")
			} else if strings.HasPrefix(part, "bg=") {
				style.Background = strings.TrimPrefix(part, "bg=")
			} else if strings.HasPrefix(part, "align=") {
				style.Align = strings.TrimPrefix(part, "align=")
			}
		}
	}

	return style
}

// StyleToString converts Style struct to string representation
func StyleToString(style Style) string {
	var parts []string

	if style.Bold {
		parts = append(parts, "bold")
	}
	if style.Italic {
		parts = append(parts, "italic")
	}
	if style.Underline {
		parts = append(parts, "underline")
	}
	if style.Color != "" {
		parts = append(parts, "color="+style.Color)
	}
	if style.Background != "" {
		parts = append(parts, "bg="+style.Background)
	}
	if style.Align != "" {
		parts = append(parts, "align="+style.Align)
	}

	return strings.Join(parts, ",")
}

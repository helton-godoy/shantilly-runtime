package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// AppConfig represents the main application configuration
type AppConfig struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
	Author  string `yaml:"author,omitempty" json:"author,omitempty"`
}

// LayoutConfig defines the overall layout structure
type LayoutConfig struct {
	Type       string      `yaml:"type" json:"type"`
	Direction  string      `yaml:"direction,omitempty" json:"direction,omitempty"`
	Components []Component `yaml:"components" json:"components"`
}

// Component represents a UI component
type Component struct {
	ID      string                 `yaml:"id,omitempty" json:"id,omitempty"`
	Type    string                 `yaml:"type" json:"type"`
	Content string                 `yaml:"content,omitempty" json:"content,omitempty"`
	Style   string                 `yaml:"style,omitempty" json:"style,omitempty"`
	Events  map[string]interface{} `yaml:"events,omitempty" json:"events,omitempty"`
	Props   map[string]interface{} `yaml:"props,omitempty" json:"props,omitempty"`
}

// Config represents the complete application configuration
type Config struct {
	App    AppConfig    `yaml:"app" json:"app"`
	Layout LayoutConfig `yaml:"layout" json:"layout"`
	// Security and runtime options will be added later
}

// LoadFromFile loads configuration from a YAML file
func LoadFromFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}

// LoadFromBytes loads configuration from byte data
func LoadFromBytes(data []byte) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}

// SaveToFile saves configuration to a YAML file
func (c *Config) SaveToFile(filename string) error {
	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file %s: %w", filename, err)
	}

	return nil
}

// Validate performs basic validation on the configuration
func (c *Config) Validate() error {
	// Validate app configuration
	if c.App.Name == "" {
		return fmt.Errorf("app.name is required")
	}
	if c.App.Version == "" {
		return fmt.Errorf("app.version is required")
	}

	// Validate layout configuration
	if c.Layout.Type == "" {
		return fmt.Errorf("layout.type is required")
	}

	// Validate layout type
	validLayoutTypes := []string{"vertical", "horizontal", "grid", "modal"}
	valid := false
	for _, t := range validLayoutTypes {
		if c.Layout.Type == t {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid layout type: %s (must be one of: %v)", c.Layout.Type, validLayoutTypes)
	}

	// Validate components
	if len(c.Layout.Components) == 0 {
		return fmt.Errorf("at least one component is required")
	}

	// Validate each component
	for i, component := range c.Layout.Components {
		if component.Type == "" {
			return fmt.Errorf("component[%d].type is required", i)
		}

		// Validate component type
		validComponentTypes := []string{"text", "button", "input", "select", "progress", "spinner", "table", "modal"}
		valid = false
		for _, t := range validComponentTypes {
			if component.Type == t {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("component[%d].type '%s' is not supported (must be one of: %v)",
				i, component.Type, validComponentTypes)
		}
	}

	return nil
}

// GetComponentByID finds a component by its ID
func (c *Config) GetComponentByID(id string) (*Component, error) {
	for _, component := range c.Layout.Components {
		if component.ID == id {
			return &component, nil
		}
	}
	return nil, fmt.Errorf("component with ID '%s' not found", id)
}

// GetComponentsByType returns all components of a specific type
func (c *Config) GetComponentsByType(componentType string) []Component {
	var components []Component
	for _, component := range c.Layout.Components {
		if component.Type == componentType {
			components = append(components, component)
		}
	}
	return components
}

// ToJSON converts configuration to JSON representation
func (c *Config) ToJSON() ([]byte, error) {
	return yaml.Marshal(c)
}

// ToYAML converts configuration to YAML representation
func (c *Config) ToYAML() ([]byte, error) {
	return yaml.Marshal(c)
}

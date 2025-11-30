package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DefaultConfig returns a default configuration for quick start
func DefaultConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:    "Shantilly Runtime App",
			Version: "1.0.0",
			Author:  "AION Framework",
		},
		Layout: LayoutConfig{
			Type:      "vertical",
			Direction: "top-to-bottom",
			Components: []Component{
				{
					Type:    "text",
					Content: "Welcome to Shantilly Runtime!",
					Style:   "bold",
				},
				{
					Type:    "button",
					ID:      "start-btn",
					Content: "Get Started",
					Events: map[string]interface{}{
						"click": map[string]interface{}{
							"type":    "run",
							"command": "echo 'Button clicked!'",
						},
					},
				},
			},
		},
	}
}

// CreateExampleConfig creates an example configuration file
func CreateExampleConfig(filename string) error {
	config := DefaultConfig()
	return config.SaveToFile(filename)
}

// ValidateFile validates a configuration file without loading it
func ValidateFile(filename string) error {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("configuration file not found: %s", filename)
	}

	// Try to load and validate
	config, err := LoadFromFile(filename)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	return config.Validate()
}

// MergeConfigurations merges multiple configurations
func MergeConfigurations(configs ...*Config) (*Config, error) {
	if len(configs) == 0 {
		return nil, fmt.Errorf("at least one configuration is required")
	}

	// Start with the first configuration
	result := configs[0]

	// Merge with subsequent configurations
	for i := 1; i < len(configs); i++ {
		config := configs[i]

		// Merge app config (non-empty fields override)
		if config.App.Name != "" {
			result.App.Name = config.App.Name
		}
		if config.App.Version != "" {
			result.App.Version = config.App.Version
		}
		if config.App.Author != "" {
			result.App.Author = config.App.Author
		}

		// Merge layout config
		if config.Layout.Type != "" {
			result.Layout.Type = config.Layout.Type
		}
		if config.Layout.Direction != "" {
			result.Layout.Direction = config.Layout.Direction
		}

		// Merge components (append)
		result.Layout.Components = append(result.Layout.Components, config.Layout.Components...)
	}

	return result, nil
}

// FindConfigFile finds configuration files in common locations
func FindConfigFile(searchPaths ...string) (string, error) {
	// Default search paths
	defaultPaths := []string{
		"shantilly.yaml",
		"shantilly.yml",
		"config.yaml",
		"config.yml",
		".shantilly.yaml",
		".shantilly.yml",
	}

	// Combine with provided paths
	allPaths := append(defaultPaths, searchPaths...)

	for _, path := range allPaths {
		// Check if path is absolute
		if filepath.IsAbs(path) {
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
			continue
		}

		// Check relative to current directory
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}

		// Check relative to current working directory
		wd, err := os.Getwd()
		if err == nil {
			fullPath := filepath.Join(wd, path)
			if _, err := os.Stat(fullPath); err == nil {
				return fullPath, nil
			}
		}
	}

	return "", fmt.Errorf("no configuration file found in search paths: %v", allPaths)
}

// GetConfigSummary returns a human-readable summary of the configuration
func (c *Config) GetConfigSummary() string {
	var summary strings.Builder

	summary.WriteString(fmt.Sprintf("📋 Configuration Summary\n"))
	summary.WriteString(fmt.Sprintf("   App: %s v%s\n", c.App.Name, c.App.Version))
	if c.App.Author != "" {
		summary.WriteString(fmt.Sprintf("   Author: %s\n", c.App.Author))
	}
	summary.WriteString(fmt.Sprintf("   Layout: %s", c.Layout.Type))
	if c.Layout.Direction != "" {
		summary.WriteString(fmt.Sprintf(" (%s)", c.Layout.Direction))
	}
	summary.WriteString(fmt.Sprintf("\n   Components: %d\n", len(c.Layout.Components)))

	// Count component types
	componentTypes := make(map[string]int)
	for _, component := range c.Layout.Components {
		componentTypes[component.Type]++
	}

	summary.WriteString("   Component Types:\n")
	for compType, count := range componentTypes {
		summary.WriteString(fmt.Sprintf("     - %s: %d\n", compType, count))
	}

	return summary.String()
}

// ValidateComponentReferences validates that all component references are valid
func (c *Config) ValidateComponentReferences() error {
	// Collect all component IDs
	componentIDs := make(map[string]bool)
	for _, component := range c.Layout.Components {
		if component.ID != "" {
			componentIDs[component.ID] = true
		}
	}

	// Check event references
	for _, component := range c.Layout.Components {
		if component.Events != nil {
			for eventType, eventAction := range component.Events {
				if actionMap, ok := eventAction.(map[string]interface{}); ok {
					// Check update_target references
					if updateTarget, ok := actionMap["update_target"].(string); ok && updateTarget != "" {
						if !componentIDs[updateTarget] {
							return fmt.Errorf("component '%s' event '%s' references unknown target '%s'",
								component.ID, eventType, updateTarget)
						}
					}
				}
			}
		}
	}

	return nil
}

// SanitizeConfig sanitizes the configuration by removing invalid or dangerous content
func (c *Config) SanitizeConfig() {
	// Sanitize component IDs
	for i := range c.Layout.Components {
		if c.Layout.Components[i].ID != "" {
			// Remove invalid characters from IDs
			c.Layout.Components[i].ID = sanitizeID(c.Layout.Components[i].ID)
		}
	}

	// Sanitize content (remove potentially dangerous content)
	for i := range c.Layout.Components {
		c.Layout.Components[i].Content = sanitizeContent(c.Layout.Components[i].Content)
	}
}

// sanitizeID removes invalid characters from component IDs
func sanitizeID(id string) string {
	// Allow only alphanumeric characters, hyphens, and underscores
	result := ""
	for _, r := range id {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			result += string(r)
		}
	}
	return result
}

// sanitizeContent removes potentially dangerous content
func sanitizeContent(content string) string {
	// Basic sanitization - remove control characters
	result := ""
	for _, r := range content {
		if r >= 32 && r != 127 { // Printable ASCII except DEL
			result += string(r)
		}
	}
	return result
}

package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.App.Name == "" {
		t.Error("Default config should have app name")
	}

	if config.App.Version == "" {
		t.Error("Default config should have app version")
	}

	if config.Layout.Type == "" {
		t.Error("Default config should have layout type")
	}

	if len(config.Layout.Components) == 0 {
		t.Error("Default config should have components")
	}
}

func TestLoadFromFile(t *testing.T) {
	// Create temporary config file
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test.yaml")

	config := DefaultConfig()
	err := config.SaveToFile(configFile)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Load the config
	loadedConfig, err := LoadFromFile(configFile)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if loadedConfig.App.Name != config.App.Name {
		t.Error("Loaded config name doesn't match")
	}

	if loadedConfig.App.Version != config.App.Version {
		t.Error("Loaded config version doesn't match")
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  DefaultConfig(),
			wantErr: false,
		},
		{
			name: "missing app name",
			config: &Config{
				App: AppConfig{Version: "1.0.0"},
				Layout: LayoutConfig{
					Type:       "vertical",
					Components: []Component{{Type: "text"}},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid layout type",
			config: &Config{
				App: AppConfig{Name: "Test", Version: "1.0.0"},
				Layout: LayoutConfig{
					Type:       "invalid",
					Components: []Component{{Type: "text"}},
				},
			},
			wantErr: true,
		},
		{
			name: "missing components",
			config: &Config{
				App: AppConfig{Name: "Test", Version: "1.0.0"},
				Layout: LayoutConfig{
					Type:       "vertical",
					Components: []Component{},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetComponentByID(t *testing.T) {
	config := &Config{
		Layout: LayoutConfig{
			Type: "vertical",
			Components: []Component{
				{ID: "btn1", Type: "button"},
				{ID: "btn2", Type: "button"},
				{Type: "text"}, // No ID
			},
		},
	}

	// Test existing component
	component, err := config.GetComponentByID("btn1")
	if err != nil {
		t.Errorf("Expected to find component btn1, got error: %v", err)
	}
	if component.ID != "btn1" {
		t.Error("Found wrong component")
	}

	// Test non-existing component
	_, err = config.GetComponentByID("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existing component")
	}
}

func TestGetComponentsByType(t *testing.T) {
	config := &Config{
		Layout: LayoutConfig{
			Type: "vertical",
			Components: []Component{
				{Type: "button"},
				{Type: "text"},
				{Type: "button"},
				{Type: "input"},
			},
		},
	}

	buttons := config.GetComponentsByType("button")
	if len(buttons) != 2 {
		t.Errorf("Expected 2 buttons, got %d", len(buttons))
	}

	texts := config.GetComponentsByType("text")
	if len(texts) != 1 {
		t.Errorf("Expected 1 text, got %d", len(texts))
	}
}

func TestValidateComponentType(t *testing.T) {
	tests := []struct {
		componentType string
		wantErr       bool
	}{
		{"text", false},
		{"button", false},
		{"input", false},
		{"invalid", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.componentType, func(t *testing.T) {
			err := ValidateComponentType(tt.componentType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateComponentType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLayoutType(t *testing.T) {
	tests := []struct {
		layoutType string
		wantErr    bool
	}{
		{"vertical", false},
		{"horizontal", false},
		{"grid", false},
		{"modal", false},
		{"invalid", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.layoutType, func(t *testing.T) {
			err := ValidateLayoutType(tt.layoutType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLayoutType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseStyleString(t *testing.T) {
	tests := []struct {
		input    string
		expected Style
	}{
		{"", Style{}},
		{"bold", Style{Bold: true}},
		{"bold,italic", Style{Bold: true, Italic: true}},
		{"color=red", Style{Color: "red"}},
		{"bg=blue", Style{Background: "blue"}},
		{"align=center", Style{Align: "center"}},
		{"bold,color=red,align=center", Style{Bold: true, Color: "red", Align: "center"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParseStyleString(tt.input)
			if result.Bold != tt.expected.Bold ||
				result.Italic != tt.expected.Italic ||
				result.Underline != tt.expected.Underline ||
				result.Color != tt.expected.Color ||
				result.Background != tt.expected.Background ||
				result.Align != tt.expected.Align {
				t.Errorf("ParseStyleString(%q) = %+v, want %+v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStyleToString(t *testing.T) {
	tests := []struct {
		style    Style
		expected string
	}{
		{Style{}, ""},
		{Style{Bold: true}, "bold"},
		{Style{Bold: true, Italic: true}, "bold,italic"},
		{Style{Color: "red"}, "color=red"},
		{Style{Background: "blue"}, "bg=blue"},
		{Style{Align: "center"}, "align=center"},
		{Style{Bold: true, Color: "red", Align: "center"}, "bold,color=red,align=center"},
	}

	for _, tt := range tests {
		t.Run("StyleToString", func(t *testing.T) {
			result := StyleToString(tt.style)
			if result != tt.expected {
				t.Errorf("StyleToString(%+v) = %q, want %q", tt.style, result, tt.expected)
			}
		})
	}
}

func TestCreateExampleConfig(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "example.yaml")

	err := CreateExampleConfig(configFile)
	if err != nil {
		t.Fatalf("CreateExampleConfig failed: %v", err)
	}

	// Check if file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Error("Example config file was not created")
	}

	// Try to load and validate
	config, err := LoadFromFile(configFile)
	if err != nil {
		t.Fatalf("Failed to load example config: %v", err)
	}

	if err := config.Validate(); err != nil {
		t.Errorf("Example config is invalid: %v", err)
	}
}

func TestFindConfigFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a config file
	configFile := filepath.Join(tmpDir, "shantilly.yaml")
	err := CreateExampleConfig(configFile)
	if err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	// Change to temp directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(originalWd)

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Test finding config file
	found, err := FindConfigFile()
	if err != nil {
		t.Errorf("FindConfigFile failed: %v", err)
	}

	// Check if found file is one of the expected names
	expectedNames := []string{"shantilly.yaml", "config.yaml"}
	foundValid := false
	for _, name := range expectedNames {
		if filepath.Base(found) == name {
			foundValid = true
			break
		}
	}

	if !foundValid {
		t.Errorf("FindConfigFile found %q, expected one of %v", found, expectedNames)
	}
}

package app

import (
	"fmt"

	"github.com/helton-godoy/shantilly-runtime/pkg/config"
	"github.com/spf13/cobra"
)

// runApp executes the TUI application
func runApp(configFile string, cmd *cobra.Command) error {
	// TODO: Implement after creating config and runtime packages
	return fmt.Errorf("runtime not implemented yet - this is a placeholder")
}

// validateConfig validates a YAML configuration file
func validateConfig(configFile string, cmd *cobra.Command) error {
	// Load configuration
	cfg, err := config.LoadFromFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	fmt.Printf("✅ Configuration file '%s' is valid\n", configFile)

	// Show configuration summary if verbose
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		fmt.Printf("%s", cfg.GetConfigSummary())
	}

	return nil
}

// showExamples displays example configurations
func showExamples() {
	examples := []struct {
		name string
		yaml string
	}{
		{
			name: "Simple Hello World",
			yaml: `app:
  name: "Hello World"
  version: "1.0.0"

layout:
  type: "vertical"
  components:
    - type: "text"
      content: "Hello, Shantilly Runtime!"
      style: "bold"
    
    - type: "button"
      label: "Click Me"
      on_click:
        type: "run"
        command: "echo 'Button clicked!'"`,
		},
		{
			name: "Form with Input",
			yaml: `app:
  name: "Input Form"
  version: "1.0.0"

layout:
  type: "vertical"
  components:
    - type: "text"
      content: "Please enter your name:"
    
    - type: "input"
      id: "name"
      placeholder: "Your name"
      required: true
    
    - type: "button"
      label: "Submit"
      on_click:
        type: "run"
        command: "echo 'Hello, {{.name}}!'"
        update_target: "name"`,
		},
		{
			name: "Interactive Menu",
			yaml: `app:
  name: "Interactive Menu"
  version: "1.0.0"

layout:
  type: "vertical"
  components:
    - type: "text"
      content: "Choose an option:"
      style: "underline"
    
    - type: "select"
      id: "choice"
      options:
        - "Option 1: Run tests"
        - "Option 2: Build project"
        - "Option 3: Deploy"
        - "Option 4: Exit"
    
    - type: "button"
      label: "Execute"
      on_click:
        type: "conditional"
        conditions:
          - when: "{{.choice}} == 'Option 1: Run tests'"
            then:
              type: "run"
              command: "go test ./..."
          - when: "{{.choice}} == 'Option 2: Build project'"
            then:
              type: "run"
              command: "go build ./..."
          - when: "{{.choice}} == 'Option 3: Deploy'"
            then:
              type: "run"
              command: "docker build -t app ."
          - when: "{{.choice}} == 'Option 4: Exit'"
            then:
              type: "exit"
              code: 0`,
		},
	}

	fmt.Println("📚 Shantilly Runtime Examples\n")
	fmt.Println("Copy and paste these examples to get started:\n")

	for i, example := range examples {
		fmt.Printf("%d. %s\n", i+1, example.name)
		fmt.Println("```yaml")
		fmt.Println(example.yaml)
		fmt.Println("```\n")
	}

	fmt.Println("💡 Save an example to a file (e.g., example.yaml) and run:")
	fmt.Println("   shantilly-runtime run example.yaml")
}

package app

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Run executes the main application
func Run(version, commit, buildDate string) error {
	rootCmd := NewRootCommand(version, commit, buildDate)
	return rootCmd.Execute()
}

// NewRootCommand creates the root CLI command
func NewRootCommand(version, commit, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shantilly-runtime",
		Short: "Runtime TUI Declarativo com orquestração AION",
		Long: `Shantilly Runtime é uma implementação limpa e focada do runtime TUI declarativo,
desenvolvida 100% pelo framework AION (AI Orchestration Native).

Use YAML para definir interfaces de terminal interativas com componentes ricos,
eventos avançados e integração com scripts shell.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, buildDate),
	}

	// Add flags
	cmd.PersistentFlags().StringP("config", "f", "", "YAML configuration file")
	cmd.PersistentFlags().StringP("output", "o", "", "Output format (json|yaml)")
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode")

	// Add subcommands
	cmd.AddCommand(NewVersionCommand(version, commit, buildDate))
	cmd.AddCommand(NewRunCommand())
	cmd.AddCommand(NewValidateCommand())
	cmd.AddCommand(NewExamplesCommand())

	return cmd
}

// NewVersionCommand creates the version command
func NewVersionCommand(version, commit, buildDate string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Shantilly Runtime %s\n", version)
			fmt.Printf("Commit: %s\n", commit)
			fmt.Printf("Built: %s\n", buildDate)
			fmt.Printf("Go: %s\n", runtime.Version())
		},
	}
}

// NewRunCommand creates the run command
func NewRunCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run [config-file]",
		Short: "Run TUI application from YAML configuration",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := ""
			if len(args) > 0 {
				configFile = args[0]
			} else {
				var err error
				configFile, err = cmd.Flags().GetString("config")
				if err != nil {
					return err
				}
			}

			if configFile == "" {
				return fmt.Errorf("configuration file is required")
			}

			return runApp(configFile, cmd)
		},
	}
}

// NewValidateCommand creates the validate command
func NewValidateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "validate [config-file]",
		Short: "Validate YAML configuration file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return validateConfig(args[0], cmd)
		},
	}
}

// NewExamplesCommand creates the examples command
func NewExamplesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "examples",
		Short: "Show example configurations",
		Run: func(cmd *cobra.Command, args []string) {
			showExamples()
		},
	}
}

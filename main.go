// wacli - A command-line interface for WhatsApp Web API
// Fork of steipete/wacli with additional features and improvements
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
	// Date is set at build time via ldflags
	Date = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wacli",
	Short: "A CLI tool for interacting with WhatsApp Web API",
	Long: `wacli is a command-line interface for sending and receiving
WhatsApp messages via the WhatsApp Web multi-device API.

It supports sending text messages, media files, and managing
WhatsApp sessions from the terminal.`,
	SilenceUsage:  true,
	SilenceErrors: true, // handle errors ourselves for cleaner output
}

// versionCmd prints the current version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wacli version %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Persistent flags available to all subcommands
	// Using ~/.wacli/config.yaml as the default config location
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file path (default: $HOME/.wacli/config.yaml)")
	// Default verbose to false; I usually run with -v only when debugging session issues
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose/debug output")
	// Shorthand -V for version to avoid conflicts with -v (verbose)
	rootCmd.Flags().BoolP("version", "V", false, "print version information and exit")
}

func main() {
	// Check for -V/--version flag before executing subcommands
	if v, _ := rootCmd.Flags().GetBool("version"); v {
		fmt.Printf("wacli version %s (commit: %s, built: %s)\n", Version, Commit, Date)
		return
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

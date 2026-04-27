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
	SilenceUsage: true,
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
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file path (default: $HOME/.wacli/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose/debug output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

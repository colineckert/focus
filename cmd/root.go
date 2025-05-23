/*
Copyright © 2025 Colin Eckert <colineckert10@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "focus",
	Short: "A simple CLI Pomodoro timer",
	Long: `Focus is a minimal command-line Pomodoro timer
to help you stay productive in timed sprints.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Focus! Use 'focus start' to begin a Pomodoro session.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

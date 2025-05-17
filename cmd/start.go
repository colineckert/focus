/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*
Example output:

$ focus start --duration 25
🍅 Focus session started: 25 minutes
[#####.............] 10m remaining
*/

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Pomodoro session",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		duration, err := cmd.Flags().GetInt("duration")
		if err != nil {
			fmt.Println("Error retrieving duration flag:", err)
			return
		}
		fmt.Printf("🍅 Focus session started: %d minutes\n", duration)

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().IntP("duration", "d", 25, "Duration of the Pomodoro session in minutes")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/colineckert/focus/internal/display"
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
	Long:  `Start a Pomodoro focus session with the option to specify the duration.`,
	Run: func(cmd *cobra.Command, args []string) {
		duration, err := cmd.Flags().GetInt("duration")
		if err != nil {
			fmt.Println("Error retrieving duration flag:", err)
			return
		}
		if duration < 1 {
			fmt.Println("Duration must be positive.")
			return
		}

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		endTime := time.Now().Add(time.Duration(duration) * time.Minute)
		fmt.Printf("🍅 Focus session started: %d minutes (ends %s)\n", duration, endTime.Format("03:04 PM"))

		for now := range ticker.C {
			remaining := endTime.Sub(now)

			if remaining <= 0 {
				fmt.Println("\n⏰ Time's up! Take a break.")
				// TODO: Update the stats to reflect the completed Pomodoro session
				break
			}

			// Print progress bar
			minutes := int(remaining.Minutes())
			seconds := int(remaining.Seconds()) % 60
			fmt.Printf("\r[%-25s] %02d:%02d remaining", display.ProgressBar(duration, remaining, 25), minutes, seconds)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntP("duration", "d", 25, "Duration of the Pomodoro session in minutes")
}

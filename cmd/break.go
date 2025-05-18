/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

/*
Example output:

$ focus break --long
☕ Long break started: 15 minutes
*/

// breakCmd represents the break command
var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Break for a short period",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		duration := 5 // Default short break duration
		long, err := cmd.Flags().GetBool("long")
		if err != nil {
			fmt.Println("Error retrieving long flag:", err)
			return
		}

		if long {
			fmt.Println("🛏️  Long break started: 15 minutes")
			duration = 15 // Long break duration
		} else {
			fmt.Println("☕ Short break started: 5 minutes")
		}

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		endTime := time.Now().Add(time.Duration(duration) * time.Minute)

		for now := range ticker.C {
			remaining := endTime.Sub(now)

			if remaining <= 0 {
				fmt.Println("\n⏰ Break's over! Time to get back to work.")
				break
			}

			// Print progress bar
			minutes := int(remaining.Minutes())
			seconds := int(remaining.Seconds()) % 60
			fmt.Printf("\r[%-25s] %02d:%02d remaining", progressBar(duration, remaining, 25), minutes, seconds)
		}
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)

	breakCmd.Flags().BoolP("long", "l", false, "Take a long break")
}

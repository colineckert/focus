/*
Copyright © 2025 Colin Eckert <colineckert10@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/colineckert/focus/internal/session"
	"github.com/spf13/cobra"
)

/*
Example output:
$ focus stats
📊 Completed today:
🍅 3 Pomodoros
☕ 2 Short Breaks
🛏️ 1 Long Break
*/

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "View the day's Pomodoro statistics",
	Long:  `See how many Pomodoros and breaks you've completed today.`,
	Run: func(cmd *cobra.Command, args []string) {
		allSessions, err := session.ReadAllSessions()
		if err != nil {
			fmt.Println("Error reading sessions:", err)
			return
		}

		stats := session.CalculateStats(allSessions)
		fmt.Printf("📊 Completed today:\n")
		fmt.Printf("🍅 %d Pomodoros\n", stats.Focus)
		fmt.Printf("☕ %d Short Breaks\n", stats.Break)
		fmt.Printf("🛏️  %d Long Breaks\n", stats.LongBreak)
		fmt.Printf("⏳ Total Time: %d minutes\n", stats.TotalTime)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	// TODO: Add flags for filtering by date or session type
}

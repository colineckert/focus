package session

import (
	"time"
)

func CalculateStats(sessions []Session) Stats {
	stats := Stats{}

	for _, session := range sessions {
		if !isToday(session.StartedAt) {
			continue
		}

		switch session.Type {
		case Focus:
			stats.Focus++
		case Break:
			stats.Break++
		case LongBreak:
			stats.LongBreak++
		}
		// consider only incrementing total time for focus sessions
		stats.TotalTime += session.DurationMinutes
	}

	return stats
}

func isToday(timestamp string) bool {
	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return false // or handle the error
	}

	now := time.Now()
	year, month, day := now.Date()
	location := now.Location()

	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, location)
	endOfDay := startOfDay.Add(24 * time.Hour)

	return parsedTime.After(startOfDay) && parsedTime.Before(endOfDay)
}

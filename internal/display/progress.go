package display

import "time"

func ProgressBar(totalDuration int, remaining time.Duration, barLength int) string {
	totalSeconds := totalDuration * 60
	remainingSeconds := int(remaining.Seconds())
	progress := (totalSeconds - remainingSeconds) * barLength / totalSeconds

	bar := ""
	for i := range barLength {
		if i < progress {
			bar += "#"
		} else {
			bar += "."
		}
	}

	return bar
}

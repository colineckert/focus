package session

/*
sessions.json example:
[
  {
    "type": "focus",
    "duration_minutes": 25,
    "started_at": "2025-05-17T14:00:00Z"
  },
  {
    "type": "break",
    "duration_minutes": 5,
    "started_at": "2025-05-17T14:30:00Z"
  }
]
*/

type SessionType string

const (
	Focus     SessionType = "focus"
	Break     SessionType = "break"
	LongBreak SessionType = "long_break"
)

// Session represents a Pomodoro focus or break session.
type Session struct {
	Type            SessionType `json:"type"`
	DurationMinutes int         `json:"duration_minutes"`
	StartedAt       string      `json:"started_at"`
}

type Stats struct {
	Focus     int
	Break     int
	LongBreak int
	TotalTime int // in minutes
}

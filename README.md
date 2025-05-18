# ⏱️ Focus CLI

A simple and customizable Pomodoro (focus timer) CLI tool written in Go.

Stay productive using the Pomodoro Technique directly from your terminal — start work sessions, take breaks, get notified when time is up, and track your focus streaks over time.

## 📦 Features

✅ MVP
- [x] Start timer (pomodoro, short break, long break).
- [x] Configurable durations (via CLI args).
- [x] Show countdown in real time or progress updates (e.g., every minute).
- [x] ASCII progress bar during timer.
- [ ] Simple logging of completed sessions (store in a local JSON/CSV file).
- [ ] View today’s stats.
- [ ] Notify user when session ends (desktop notification, sound, or terminal output).

🛠️ Stretch Features
- [ ] CLI args or interactive prompts.
- [ ] Streak tracking.
- [ ] TUI mode (e.g., live countdown in a curses UI).
- [ ] Idle detection or pause/resume (advanced).
- [ ] GitHub-style heatmap of productivity (using focus stats --week).

## 🚀 Installation

Clone and build:

```bash
git clone https://github.com/colineckert/focus
cd focus
go build -o focus
./focus --help
```

package session

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
Record a session to disk in ~/.focus/sessions.json.
*/
func WriteSessionToJSON(session Session) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Create .focus directory if it doesn't exist
	focusDir := filepath.Join(homeDir, ".focus")
	if err := os.MkdirAll(focusDir, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(focusDir, "sessions.json")

	var sessions []Session

	// try reading ~/.focus/sessions.json if it exists
	if data, err := os.ReadFile(filePath); err == nil && len(data) > 0 {
		// unmarshal into a slice of Session
		if err := json.Unmarshal(data, &sessions); err != nil {
			return fmt.Errorf("failed to parse existing sessions.json: %w", err)
		}
	}

	// Append the new session
	sessions = append(sessions, session)

	// Marshal the updated sessions slice to JSON
	newData, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the file
	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		return err
	}

	return nil
}

/*
Return a slice of Session structs from disk.
If the file doesn't exist or is empty, return an empty slice (not an error).
*/
func ReadAllSessions() ([]Session, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(homeDir, ".focus", "sessions.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Session{}, nil // File doesn't exist, return empty slice
		}
		return nil, err // Some other error occurred
	}

	if len(data) == 0 {
		return []Session{}, nil // File is empty, return empty slice
	}

	var sessions []Session
	if err := json.Unmarshal(data, &sessions); err != nil {
		return nil, fmt.Errorf("failed to parse sessions.json: %w", err)
	}

	return sessions, nil
}

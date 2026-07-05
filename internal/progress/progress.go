// Package progress persists a player's advancement through VimHero:
// which day is unlocked, best keystroke count per challenge, and the
// daily play streak.
package progress

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type ChallengeResult struct {
	Cleared  bool `json:"cleared"`
	BestKeys int  `json:"best_keys"`
	Stars    int  `json:"stars"`
}

type State struct {
	UnlockedDay int                        `json:"unlocked_day"`
	Results     map[string]ChallengeResult `json:"results"` // key: "day-challengeIndex"
	Streak      int                        `json:"streak"`
	LastPlayed  string                     `json:"last_played"` // YYYY-MM-DD
}

func defaultState() *State {
	return &State{
		UnlockedDay: 1,
		Results:     map[string]ChallengeResult{},
	}
}

func filePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".vimhero")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "progress.json"), nil
}

func Load() *State {
	path, err := filePath()
	if err != nil {
		return defaultState()
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return defaultState()
	}
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return defaultState()
	}
	if s.Results == nil {
		s.Results = map[string]ChallengeResult{}
	}
	if s.UnlockedDay < 1 {
		s.UnlockedDay = 1
	}
	return &s
}

func (s *State) Save() error {
	path, err := filePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func key(day, challenge int) string {
	return itoa(day) + "-" + itoa(challenge)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

// RecordClear stores a challenge clear if it improves on any prior best,
// and unlocks the next day when every challenge for `day` is cleared.
func (s *State) RecordClear(day, challenge, keystrokes, stars, totalChallengesInDay int) {
	k := key(day, challenge)
	prev, ok := s.Results[k]
	if !ok || keystrokes < prev.BestKeys {
		s.Results[k] = ChallengeResult{Cleared: true, BestKeys: keystrokes, Stars: stars}
	} else {
		prev.Cleared = true
		if stars > prev.Stars {
			prev.Stars = stars
		}
		s.Results[k] = prev
	}

	allCleared := true
	for i := 0; i < totalChallengesInDay; i++ {
		if !s.Results[key(day, i)].Cleared {
			allCleared = false
			break
		}
	}
	if allCleared && s.UnlockedDay <= day {
		s.UnlockedDay = day + 1
	}

	s.bumpStreak()
}

func (s *State) bumpStreak() {
	today := time.Now().Format("2006-01-02")
	if s.LastPlayed == today {
		return
	}
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	if s.LastPlayed == yesterday {
		s.Streak++
	} else {
		s.Streak = 1
	}
	s.LastPlayed = today
}

func (s *State) ChallengeResult(day, challenge int) (ChallengeResult, bool) {
	r, ok := s.Results[key(day, challenge)]
	return r, ok
}

func (s *State) IsDayUnlocked(day int) bool {
	return day <= s.UnlockedDay
}

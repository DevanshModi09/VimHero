// Package curriculum holds the static "zero to hero in 45 days" Vim
// lesson plan: for each day, a short lesson and one or more challenges
// that are played on the editor engine.
package curriculum

type Kind int

const (
	// KindGoal: move the cursor onto GoalPos using only motions; the
	// buffer content itself never changes.
	KindGoal Kind = iota
	// KindEdit: transform Start into exactly Target (line for line).
	KindEdit
)

type Pos struct {
	Row, Col int
}

type Challenge struct {
	Title        string
	Instructions string // explains what the new command(s) actually do, then what to do
	Tip          string // a shown-on-screen tip: a deeper detail or related trick
	NewKeys      []string
	Start        []string
	CursorStart  Pos
	Kind         Kind
	GoalPos      Pos
	Target       []string
	Par          int // keystrokes for a 3-star clear
}

type Day struct {
	Number     int
	Week       string
	Title      string
	Summary    string
	Challenges []Challenge
}

// Stars converts a keystroke count into a 1-3 star rating against a
// challenge's par. 3 stars at or under par, 2 within double par, else 1.
func Stars(keystrokes, par int) int {
	if par <= 0 {
		return 3
	}
	switch {
	case keystrokes <= par:
		return 3
	case keystrokes <= par*2:
		return 2
	default:
		return 1
	}
}

// All returns the full 45-day curriculum in order.
func All() []Day {
	return days
}

func ByNumber(n int) (Day, bool) {
	for _, d := range days {
		if d.Number == n {
			return d, true
		}
	}
	return Day{}, false
}

func Total() int {
	return len(days)
}

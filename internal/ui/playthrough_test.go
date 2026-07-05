package ui

import "testing"

// TestCurriculumSolvable scripts the intended solution for every
// authored challenge and verifies it actually reaches the win condition
// against the real editor engine. This catches off-by-one errors in
// GoalPos/Target/CursorStart before a player ever hits them.
func TestCurriculumSolvable(t *testing.T) {
	cases := []struct {
		day, challenge int
		keys           []string
	}{
		{1, 0, []string{"j", "j", "l", "l"}},
		{1, 1, []string{"j", "j", "j", "j", "j", "l", "l", "l", "l", "l", "l"}},

		{2, 0, []string{"w", "w", "w"}},
		{2, 1, []string{"b", "b", "b"}},
		{2, 2, []string{"e", "e", "e"}},

		{3, 0, []string{"i", "h", "e", "l", "l", "o", " ", "v", "i", "m", "esc"}},
		{3, 1, []string{"A", " ", "w", "o", "r", "l", "d", "esc"}},
		{3, 2, []string{"o", "s", "e", "c", "o", "n", "d", " ", "l", "i", "n", "e", "esc"}},

		{4, 0, []string{"g", "g"}},
		{4, 1, []string{"G"}},
		{4, 2, []string{"x"}},

		{5, 0, []string{"d", "d"}},
		{5, 1, []string{"y", "y", "p"}},
		{5, 2, []string{"y", "y", "P"}},
	}

	m := NewModel()
	for _, tc := range cases {
		dayIdx := tc.day - 1
		day := m.days[dayIdx]
		mm := m.enterChallenge(dayIdx, tc.challenge)
		for _, k := range tc.keys {
			mm.buf.Input(k)
		}
		if !mm.checkWin() {
			t.Errorf("day %d challenge %d (%s) not solved by scripted keys %v\n  got lines: %q\n  got cursor: %+v\n  want: kind=%v target=%v goal=%+v",
				tc.day, tc.challenge, day.Challenges[tc.challenge].Title, tc.keys,
				mm.buf.Lines, mm.buf.Cursor,
				mm.challenge.Kind, mm.challenge.Target, mm.challenge.GoalPos)
		}
	}
}

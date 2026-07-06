package ui

import "testing"

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

		{6, 0, []string{"d", "w"}},
		{6, 1, []string{"c", "w", "q", "u", "i", "c", "k", "esc"}},
		{6, 2, []string{"c", "i", "w", "w", "o", "r", "d", "esc"}},

		{7, 0, []string{"d", "a", "w"}},
		{7, 1, []string{"D"}},
		{7, 2, []string{"C", "y", "o", "u", "esc"}},

		{8, 0, []string{"c", "w", "q", "u", "i", "c", "k", "esc", "j", "w", "c", "w", "o", "v", "e", "r", "esc"}},
		{8, 1, []string{"d", "d", "g", "g", "P"}},
		{8, 2, []string{"A", " ", "w", "o", "r", "l", "d", "esc", "G", "o", "S", "e", "e", " ", "y", "o", "u", " ", "s", "o", "o", "n", "esc"}},
		{8, 3, []string{"c", "w", "q", "u", "i", "c", "k", "esc", "w", "w", "w", "w", "c", "w", "o", "v", "e", "r", "esc"}},
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

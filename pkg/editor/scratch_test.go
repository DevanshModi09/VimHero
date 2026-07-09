package editor

import "testing"

func TestScratchDay13(t *testing.T) {
	run := func(name string, lines []string, cursor Pos, keys []string) {
		b := New(append([]string(nil), lines...), cursor)
		for _, k := range keys {
			b.Input(k)
		}
		t.Logf("%s => lines=%q cursor=%+v", name, b.Lines, b.Cursor)
	}

	// Ch1: teach 0 — jump to absolute col 0 on an indented line.
	run("ch1_0", []string{"    indented text here"}, Pos{0, 15}, []string{"0"})

	// Ch2: teach ^ — jump to first non-blank char.
	run("ch2_caret", []string{"    the real text starts here"}, Pos{0, 28}, []string{"^"})

	// Ch3: teach $ with a count — jump to end of the 3rd line down.
	run("ch3_dollar_count", []string{"short", "medium line here", "the longest line of them all"}, Pos{0, 0}, []string{"3", "$"})

	// Ch4: combine operators with these motions — d$ and c^
	run("ch4_ddollar", []string{"keep this part but not this trailing junk"}, Pos{0, 14}, []string{"d", "$"})
	run("ch4_ccaret", []string{"    XXXX the rest of this line stays"}, Pos{0, 36}, []string{"c", "^"})

	// Capstone ideas
	run("cap_line1", []string{"    fix the indentation on this line"}, Pos{0, 20}, []string{"^", "d", "0"})
	run("cap_line2", []string{"trim the trailing mess right here###"}, Pos{0, 0}, []string{"$"})
}

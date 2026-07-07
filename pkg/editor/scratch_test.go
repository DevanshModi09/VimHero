package editor

import (
	"fmt"
	"testing"
)

func typeKeys(s string) []string {
	out := make([]string, 0, len(s))
	for _, r := range s {
		out = append(out, string(r))
	}
	return out
}

func run(b *Buffer, keys []string) {
	for _, k := range keys {
		b.Input(k)
	}
}

func TestScratchDay10(t *testing.T) {
	// Challenge 1: single s, replace 1 char with longer text
	b := New([]string{"I have 5 dogs"}, Pos{0, 0})
	keys := append([]string{"w", "w", "s"}, append(typeKeys("five"), "esc")...)
	run(b, keys)
	fmt.Printf("C1 lines=%q cursor=%+v keys=%d\n", b.Lines, b.Cursor, len(keys))

	// Challenge 2: count-prefixed s, replace 3 chars with new text
	b2 := New([]string{"The CPU is 100% busy"}, Pos{0, 0})
	keys2 := append([]string{"w", "w", "w", "3", "s"}, append(typeKeys("50"), "esc")...)
	run(b2, keys2)
	fmt.Printf("C2 lines=%q cursor=%+v keys=%d\n", b2.Lines, b2.Cursor, len(keys2))

	// Challenge 3: S, replace whole line
	b3 := New([]string{"this line is a mess and wrong"}, Pos{0, 0})
	keys3 := append([]string{"S"}, append(typeKeys("this line is correct now"), "esc")...)
	run(b3, keys3)
	fmt.Printf("C3 lines=%q cursor=%+v keys=%d\n", b3.Lines, b3.Cursor, len(keys3))

	// Challenge 4: count-prefixed S, merge 2 lines into 1
	b4 := New([]string{"draft line one", "draft line two"}, Pos{0, 0})
	keys4 := append([]string{"2", "S"}, append(typeKeys("final combined line"), "esc")...)
	run(b4, keys4)
	fmt.Printf("C4 lines=%q cursor=%+v keys=%d\n", b4.Lines, b4.Cursor, len(keys4))

	// Challenge 5: capstone - count-s, S, count-s across 3 lines
	b5 := New([]string{
		"the score is 42 points",
		"this whole line is nonsense and needs replacing",
		"fix typo1 in here",
	}, Pos{0, 0})
	keys5 := []string{}
	keys5 = append(keys5, "w", "w", "w")
	keys5 = append(keys5, "2", "s")
	keys5 = append(keys5, typeKeys("one hundred")...)
	keys5 = append(keys5, "esc")
	keys5 = append(keys5, "j")
	keys5 = append(keys5, "S")
	keys5 = append(keys5, typeKeys("clean sentence here")...)
	keys5 = append(keys5, "esc")
	keys5 = append(keys5, "j", "w")
	keys5 = append(keys5, "5", "s")
	keys5 = append(keys5, typeKeys("this")...)
	keys5 = append(keys5, "esc")
	run(b5, keys5)
	fmt.Printf("C5 lines=%q cursor=%+v keys=%d\n", b5.Lines, b5.Cursor, len(keys5))
}

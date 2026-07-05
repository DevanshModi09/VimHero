package curriculum

// days holds the full curriculum. Only Days 1-5 (Week 1: Basic Movement
// & Modes) are authored so far — Days 6-45 are added incrementally.
var days = []Day{
	{
		Number: 1,
		Week:   "Week 1: Basic Movement",
		Title:  "hjkl — Move Without Arrow Keys",
		Summary: "Vim keeps your fingers on the home row: h moves left, l moves right, " +
			"j moves down (think: j looks like a hook dropping down), k moves up. " +
			"No mouse, no arrow keys — just h j k l.",
		Challenges: []Challenge{
			{
				Title: "First Steps",
				Instructions: "h moves the cursor one cell left, l moves it one cell right, " +
					"j moves it one cell down, k moves it one cell up. There's no diagonal " +
					"key — to move diagonally you just press two of them, like j then l. " +
					"Reach the G using only h j k l.",
				Tip: "Tip: h j k l sit right under your fingers on the home row (no need to move " +
					"your hand to an arrow-key cluster) — that's the whole reason Vim uses them.",
				NewKeys: []string{"h", "j", "k", "l"},
				Start: []string{
					".....",
					".....",
					"..G..",
					".....",
					".....",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{2, 2},
				Par:         4,
			},
			{
				Title: "Bigger Maze",
				Instructions: "Same four keys, longer trip: chain together presses of h j k l — " +
					"in any order — to cross the whole grid and land exactly on the G.",
				Tip: "Tip: you can prefix a motion with a number to repeat it, e.g. 5j moves down " +
					"5 cells in a single keystroke instead of pressing j five times — try it here.",
				NewKeys: []string{"h", "j", "k", "l"},
				Start: []string{
					".........",
					".........",
					".........",
					".........",
					".........",
					"......G..",
					".........",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{5, 6},
				Par:         11,
			},
		},
	},
	{
		Number: 2,
		Week:   "Week 1: Basic Movement",
		Title:  "Word Motions — w, e, b",
		Summary: "Moving one letter at a time is slow. w jumps to the start of the next word, " +
			"e jumps to the end of the current/next word, b jumps back to the start of the " +
			"previous word. This is how fast Vim users cross a line in a few keystrokes.",
		Challenges: []Challenge{
			{
				Title: "Jump Forward",
				Instructions: "w moves the cursor to the very first character of the next word — " +
					"punctuation and symbols like $ count as words of their own. Press w " +
					"repeatedly to land exactly on the $ below.",
				Tip: "Tip: try it mentally on \"one two three\" — from \"one\", a single w lands " +
					"you on \"two\", another w lands you on \"three\".",
				NewKeys:     []string{"w"},
				Start:       []string{"jump over the $ to win"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 14},
				Par:         3,
			},
			{
				Title: "Jump Backward",
				Instructions: "b moves the cursor back to the start of a word. If the cursor " +
					"isn't already sitting on a word's first letter, the first b snaps back to " +
					"the start of the word you're currently inside — not the previous word. " +
					"Starting at the end of the line, use b to walk backward onto the $.",
				Tip: "Tip: this trips a lot of people up — from the middle of \"high\", the first " +
					"b lands on the \"h\" of \"high\" itself, and only the next b moves further back.",
				NewKeys:     []string{"b"},
				Start:       []string{"quick brown $ jumps high"},
				CursorStart: Pos{0, 23},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 12},
				Par:         3,
			},
			{
				Title: "Land On the Last Letter",
				Instructions: "e moves the cursor to the END of the current or next word, instead " +
					"of the start like w does. Use e to land precisely on the last letter of " +
					"\"target\", the third word below.",
				Tip: "Tip: w vs e is the key distinction to remember — w lands on a word's " +
					"first letter, e lands on its last. Whichever is closer costs fewer presses.",
				NewKeys:     []string{"e"},
				Start:       []string{"reach the target now"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 15},
				Par:         3,
			},
		},
	},
	{
		Number: 3,
		Week:   "Week 1: Basic Movement & Modes",
		Title:  "Insert Mode — i, A, o",
		Summary: "Normal mode is for moving and editing; Insert mode is for typing. " +
			"i enters Insert mode before the cursor, A jumps to end-of-line and enters " +
			"Insert mode, o opens a new line below and enters Insert mode. Esc always " +
			"returns you to Normal mode — you'll press it constantly.",
		Challenges: []Challenge{
			{
				Title: "Type Something",
				Instructions: "i switches from Normal mode into Insert mode, without moving the " +
					"cursor — anything you type from here on is inserted as regular text, just " +
					"like a normal text editor. Press i, type: hello vim — then press Esc to " +
					"drop back into Normal mode.",
				Tip: "Tip: Esc is one of the most-pressed keys in Vim — a lot of people remap " +
					"Caps Lock to Esc on their keyboard since it's easier to reach.",
				NewKeys:     []string{"i", "esc"},
				Start:       []string{""},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"hello vim"},
				Par:         11,
			},
			{
				Title: "Append At The End",
				Instructions: "A jumps straight to the end of the line and enters Insert mode there " +
					"— a shortcut for pressing $ and then a. No matter where your cursor " +
					"currently sits on the line, use A to append \" world\" to the end of it.",
				Tip: "Tip: lowercase a inserts right after the cursor's current spot; capital A " +
					"always jumps to end-of-line first, regardless of cursor position.",
				NewKeys:     []string{"A"},
				Start:       []string{"hello"},
				CursorStart: Pos{0, 2},
				Kind:        KindEdit,
				Target:      []string{"hello world"},
				Par:         8,
			},
			{
				Title: "Open A New Line",
				Instructions: "o opens a brand-new, empty line directly below the cursor's current " +
					"line and drops you into Insert mode on it — handy for adding a line without " +
					"first navigating to end-of-line. Use it to insert a line reading: second " +
					"line — between the two lines below.",
				Tip: "Tip: capital O does the same thing but opens the new line ABOVE the " +
					"cursor instead of below — useful when you need to insert something before " +
					"the current line.",
				NewKeys:     []string{"o"},
				Start:       []string{"first line", "third line"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"first line", "second line", "third line"},
				Par:         13,
			},
		},
	},
	{
		Number: 4,
		Week:   "Week 1: Basic Movement",
		Title:  "gg, G, and x",
		Summary: "gg jumps to the very first line, G jumps to the very last line — instant " +
			"travel across a whole file. x deletes the single character under the cursor, " +
			"your first real edit command.",
		Challenges: []Challenge{
			{
				Title: "Jump To The Top",
				Instructions: "gg (tap g twice) jumps straight to line 1 of the file, no matter " +
					"how far down you are — instant travel instead of holding k. Use it to get " +
					"from the bottom of this list back to the top.",
				Tip:         "Tip: gg accepts a count too — 5gg jumps straight to line 5, exactly.",
				NewKeys:     []string{"gg"},
				Start:       []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6"},
				CursorStart: Pos{5, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 0},
				Par:         2,
			},
			{
				Title: "Jump To The Bottom",
				Instructions: "G jumps straight to the very last line of the file — the mirror " +
					"image of gg. Use it to go from the top all the way to the bottom in a " +
					"single keystroke.",
				Tip: "Tip: like gg, G takes a count — 3G jumps to exactly line 3, same as " +
					"typing :3 and pressing enter in command mode.",
				NewKeys:     []string{"G"},
				Start:       []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{5, 0},
				Par:         1,
			},
			{
				Title: "Delete The Typo",
				Instructions: "x deletes exactly one character — whichever one is currently under " +
					"the cursor — making it your fastest fix for a stray typo. Delete the extra " +
					"l in \"helllo\" below so it reads \"hello\".",
				Tip: "Tip: capital X deletes the character BEFORE the cursor instead of under " +
					"it — reach for X when you've typed one character too many.",
				NewKeys:     []string{"x"},
				Start:       []string{"helllo"},
				CursorStart: Pos{0, 4},
				Kind:        KindEdit,
				Target:      []string{"hello"},
				Par:         1,
			},
		},
	},
	{
		Number: 5,
		Week:   "Week 1: Basic Movement & Modes",
		Title:  "dd, yy, p, P — Delete, Yank, Paste",
		Summary: "dd deletes (and stores) the whole current line. yy yanks (copies) it " +
			"without deleting. p pastes after the current line, P pastes before it. " +
			"Together these are how you move and duplicate lines without ever touching a mouse.",
		Challenges: []Challenge{
			{
				Title: "Delete The Line",
				Instructions: "dd deletes the entire current line in one shot — wherever the " +
					"cursor sits on that line, the whole line goes, and it's saved so you could " +
					"paste it elsewhere. Remove the line that shouldn't be there.",
				Tip: "Tip: 3dd deletes 3 lines starting from the cursor — operators like d " +
					"combine with counts exactly like motions do.",
				NewKeys:     []string{"dd"},
				Start:       []string{"keep this", "DELETE THIS LINE", "keep this too"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"keep this", "keep this too"},
				Par:         2,
			},
			{
				Title: "Duplicate Below",
				Instructions: "yy yanks (copies) the current line without deleting it, and p pastes " +
					"whatever was last yanked or deleted directly below the cursor's line. Yank " +
					"the first line, then paste a copy of it right below itself.",
				Tip: "Tip: think of y as Vim's \"copy\" and d as its \"cut\" — both fill the same " +
					"register, and p is always the paste that follows either one.",
				NewKeys:     []string{"yy", "p"},
				Start:       []string{"alpha", "beta"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"alpha", "alpha", "beta"},
				Par:         3,
			},
			{
				Title: "Paste Above",
				Instructions: "P (capital) pastes above the cursor's line instead of below it — " +
					"the only difference from lowercase p. Yank \"two\" below, then use P to " +
					"paste a copy of it above itself.",
				Tip: "Tip: lowercase p = paste after/below, capital P = paste before/above. " +
					"This same after-vs-before pattern (a/A, i/I, p/P) repeats all over Vim.",
				NewKeys:     []string{"P"},
				Start:       []string{"one", "two"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"one", "two", "two"},
				Par:         3,
			},
		},
	},
}

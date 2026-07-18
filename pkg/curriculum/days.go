package curriculum

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
					".......",
					".......",
					".......",
					"....G..",
					".......",
					".......",
					".......",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{3, 4},
				Par:         7,
			},
			{
				Title: "Bigger Maze",
				Instructions: "Same four keys, longer trip: chain together presses of h j k l — " +
					"in any order — to cross the whole grid and land exactly on the G.",
				Tip: "Tip: you can prefix a motion with a number to repeat it, e.g. 6j moves down " +
					"6 cells in a single keystroke instead of pressing j six times — try it here.",
				NewKeys: []string{"h", "j", "k", "l"},
				Start: []string{
					"..........",
					"..........",
					"..........",
					"..........",
					"..........",
					"..........",
					".......G..",
					"..........",
					"..........",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{6, 7},
				Par:         13,
			},
			{
				Title: "Zigzag Descent",
				Instructions: "The grid is bigger again — this one takes more presses than either " +
					"of the first two combined. Cross it one cell at a time with h j k l to " +
					"land on the G.",
				Tip: "Tip: there's no wrong order to mix j and l here — go all the way down first, " +
					"then across, or interleave them, the destination is the same either way.",
				Start: []string{
					"............",
					"............",
					"............",
					"............",
					"............",
					"............",
					"............",
					"..........G.",
					"............",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{7, 10},
				Par:         17,
			},
			{
				Title: "Up And Over",
				Instructions: "So far you've only needed j and l. This time the G sits above and " +
					"to the left of your start — you'll need k and h instead, moving up and " +
					"left back toward the corner.",
				Tip: "Tip: h and l share the same row of the keyboard as j and k — once all " +
					"four feel natural, you never have to look down at your hands to move.",
				Start: []string{
					"........",
					".G......",
					"........",
					"........",
					"........",
					"........",
					"........",
					"........",
				},
				CursorStart: Pos{6, 6},
				Kind:        KindGoal,
				GoalPos:     Pos{1, 1},
				Par:         10,
			},
			{
				Title: "Full Traverse",
				Instructions: "The longest hjkl trip yet: start at the bottom-left corner and " +
					"climb all the way to the G in the top-right corner, crossing the full " +
					"height and width of the grid.",
				Tip: "Tip: this is exactly the kind of long, repetitive trip a count prefix " +
					"was built for — 9k then 13l gets you there in two keystrokes total instead " +
					"of twenty-two.",
				Start: []string{
					".............G",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
				},
				CursorStart: Pos{9, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 13},
				Par:         22,
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
					"punctuation and symbols like $ count as words of their own. This line has " +
					"plenty of words to cross — press w repeatedly to land exactly on the $ at " +
					"the end.",
				Tip: "Tip: try it mentally on \"one two three\" — from \"one\", a single w lands " +
					"you on \"two\", another w lands you on \"three\".",
				NewKeys:     []string{"w"},
				Start:       []string{"jump over several small words to reach the $ sign"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 43},
				Par:         8,
			},
			{
				Title: "Jump Backward",
				Instructions: "b moves the cursor back to the start of a word. If the cursor " +
					"isn't already sitting on a word's first letter, the first b snaps back to " +
					"the start of the word you're currently inside — not the previous word. " +
					"Starting at the end of the line, use b to walk all the way back onto \"over\".",
				Tip: "Tip: this trips a lot of people up — from the middle of \"bank\", the first " +
					"b lands on the \"b\" of \"bank\" itself, and only the next b moves further back.",
				NewKeys:     []string{"b"},
				Start:       []string{"quick brown fox jumps over lazy dog near river bank"},
				CursorStart: Pos{0, 50},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 22},
				Par:         6,
			},
			{
				Title: "Land On the Last Letter",
				Instructions: "e moves the cursor to the END of the current or next word, instead " +
					"of the start like w does. Use e to hop word by word all the way to the last " +
					"letter of \"field\", the final word below.",
				Tip: "Tip: w vs e is the key distinction to remember — w lands on a word's " +
					"first letter, e lands on its last. Whichever is closer costs fewer presses.",
				NewKeys:     []string{"e"},
				Start:       []string{"reach the distant target across this wide open field"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 51},
				Par:         9,
			},
			{
				Title: "Word Hopscotch",
				Instructions: "w and e work exactly the same whether you start at the beginning " +
					"of a line or somewhere in the middle. Starting from \"four\" below, hop " +
					"forward word by word until you land on \"twelve\".",
				Tip: "Tip: word motions never care where on the line you started — only where " +
					"the word boundaries are from that point on.",
				Start:       []string{"one two three four five six seven eight nine ten eleven twelve"},
				CursorStart: Pos{0, 14},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 56},
				Par:         8,
			},
			{
				Title: "Backtrack Marathon",
				Instructions: "The longest word-motion trip yet. Starting at the very end of the " +
					"line (inside the last word), use b to walk all the way back to \"market\", " +
					"the very first word.",
				Tip: "Tip: remember the quirk from \"Jump Backward\" — your first b here snaps " +
					"to the start of the word you're already inside, not the one before it.",
				Start:       []string{"market square filled with fresh apples ripe pears sweet cherries golden plums"},
				CursorStart: Pos{0, 76},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 0},
				Par:         12,
			},
		},
	},
	{
		Number: 3,
		Week:   "Week 1: Basic Movement & Modes",
		Title:  "Insert Mode — i, A, o, O",
		Summary: "Normal mode is for moving and editing; Insert mode is for typing. " +
			"i enters Insert mode before the cursor, A jumps to end-of-line and enters " +
			"Insert mode, o opens a new line below and enters Insert mode, O opens one " +
			"above. Esc always returns you to Normal mode — you'll press it constantly.",
		Challenges: []Challenge{
			{
				Title: "Type Something",
				Instructions: "i switches from Normal mode into Insert mode, without moving the " +
					"cursor — anything you type from here on is inserted as regular text, just " +
					"like a normal text editor. Press i, type: welcome to vim insert mode — then " +
					"press Esc to drop back into Normal mode.",
				Tip: "Tip: Esc is one of the most-pressed keys in Vim — a lot of people remap " +
					"Caps Lock to Esc on their keyboard since it's easier to reach.",
				NewKeys:     []string{"i", "esc"},
				Start:       []string{""},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"welcome to vim insert mode"},
				Par:         28,
			},
			{
				Title: "Append At The End",
				Instructions: "A jumps straight to the end of the line and enters Insert mode there " +
					"— a shortcut for pressing $ and then a. No matter where your cursor " +
					"currently sits on the line, use A to append \" world of vim editing\" to " +
					"the end of it.",
				Tip: "Tip: lowercase a inserts right after the cursor's current spot; capital A " +
					"always jumps to end-of-line first, regardless of cursor position.",
				NewKeys:     []string{"A"},
				Start:       []string{"hello"},
				CursorStart: Pos{0, 2},
				Kind:        KindEdit,
				Target:      []string{"hello world of vim editing"},
				Par:         23,
			},
			{
				Title: "Open A New Line",
				Instructions: "o opens a brand-new, empty line directly below the cursor's current " +
					"line and drops you into Insert mode on it — handy for adding a line without " +
					"first navigating to end-of-line. Use it to insert a line reading: second " +
					"line goes right here — between the two lines below.",
				Tip: "Tip: capital O does the same thing but opens the new line ABOVE the " +
					"cursor instead of below — you'll use it in the very next challenge.",
				NewKeys:     []string{"o"},
				Start:       []string{"first line", "third line"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"first line", "second line goes right here", "third line"},
				Par:         29,
			},
			{
				Title: "Insert Above",
				Instructions: "O (capital) opens a new line ABOVE the cursor's current line, then " +
					"drops you into Insert mode on it — the mirror image of lowercase o. This " +
					"file is missing its opening line; use O to add \"first line\" above what's " +
					"already here.",
				Tip: "Tip: think of the a/A and i/I family — lowercase acts relative to the " +
					"cursor, capital jumps to an edge first. o/O follows the same before/after " +
					"pattern, just one line up or down instead of left or right.",
				NewKeys:     []string{"O"},
				Start:       []string{"second line", "third line"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"first line", "second line", "third line"},
				Par:         12,
			},
			{
				Title: "Sandwich The Line",
				Instructions: "Now combine all four: O to add a line above, A to append onto the " +
					"existing line, and o to add a line below — turning one lonely line into a " +
					"full three-course meal. Build: appetizer / main course dish / dessert.",
				Tip: "Tip: after O...Esc you land back on the line you just typed, not the " +
					"original one — you'll need to move down before appending to it.",
				Start:       []string{"main course"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"appetizer", "main course dish", "dessert"},
				Par:         28,
			},
		},
	},
	{
		Number: 4,
		Week:   "Week 1: Basic Movement",
		Title:  "gg, G, and x",
		Summary: "gg jumps to the very first line, G jumps to the very last line — instant " +
			"travel across a whole file, and a count prefix like 7gg jumps straight to line " +
			"7. x deletes the single character under the cursor, your first real edit command.",
		Challenges: []Challenge{
			{
				Title: "Jump To The Top",
				Instructions: "gg (tap g twice) jumps straight to line 1 of the file, no matter " +
					"how far down you are — instant travel instead of holding k. Use it to get " +
					"from the bottom of this list back to the top.",
				Tip:         "Tip: gg accepts a count too — 5gg jumps straight to line 5, exactly.",
				NewKeys:     []string{"gg"},
				Start:       []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				CursorStart: Pos{9, 0},
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
				Start:       []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{9, 0},
				Par:         1,
			},
			{
				Title: "Delete The Typo",
				Instructions: "x deletes exactly one character — whichever one is currently under " +
					"the cursor — making it your fastest fix for a stray typo. All three lines " +
					"below start with a stray extra letter — delete each one so the lines read " +
					"hello / world / test.",
				Tip: "Tip: capital X deletes the character BEFORE the cursor instead of under " +
					"it — reach for X when you've typed one character too many.",
				NewKeys:     []string{"x"},
				Start:       []string{"xhello", "xworld", "xtest"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"hello", "world", "test"},
				Par:         5,
			},
			{
				Title: "Count Your Jumps",
				Instructions: "A number before gg or G jumps straight to that exact line — no " +
					"counting keystrokes with j needed. Use a single count-prefixed jump to land " +
					"on \"seven\", the 7th line of this list.",
				Tip: "Tip: this is the same trick as 5j from Day 1's tip, applied to gg/G — " +
					"any motion that repeats can usually take a count prefix.",
				Start:       []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{6, 0},
				Par:         3,
			},
			{
				Title: "Typo Hunt",
				Instructions: "Three lines in this longer file each have a stray letter tacked " +
					"onto the front. Use gg, G, j, and k to travel between them and x to delete " +
					"each one — you decide the fastest route.",
				Tip: "Tip: G to jump straight to the bottom typo, then work your way up with k, " +
					"is often faster than counting j's down from the top.",
				Start: []string{
					"first", "xsecond", "third", "fourth", "xfifth", "sixth", "seventh", "xeighth",
				},
				CursorStart: Pos{3, 0},
				Kind:        KindEdit,
				Target: []string{
					"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth",
				},
				Par: 9,
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
					"paste it elsewhere. Two lines below don't belong — navigate to each and " +
					"remove it.",
				Tip: "Tip: 3dd deletes 3 lines starting from the cursor — operators like d " +
					"combine with counts exactly like motions do.",
				NewKeys:     []string{"dd"},
				Start:       []string{"keep1", "BAD1", "keep2", "keep3", "BAD2", "keep4"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"keep1", "keep2", "keep3", "keep4"},
				Par:         7,
			},
			{
				Title: "Duplicate Below",
				Instructions: "yy yanks (copies) the current line without deleting it, and p pastes " +
					"whatever was last yanked or deleted directly below the cursor's line — " +
					"pressing p again pastes the SAME copy a second time. Yank the first line, " +
					"then paste it twice so it appears three times total.",
				Tip: "Tip: think of y as Vim's \"copy\" and d as its \"cut\" — both fill the same " +
					"register, and p is always the paste that follows either one, as many times " +
					"as you press it.",
				NewKeys:     []string{"yy", "p"},
				Start:       []string{"template", "other1", "other2"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"template", "template", "template", "other1", "other2"},
				Par:         5,
			},
			{
				Title: "Paste Above",
				Instructions: "P (capital) pastes above the cursor's line instead of below it — " +
					"the only difference from lowercase p. Yank \"target\" below, then press P " +
					"twice to stack two extra copies above the original.",
				Tip: "Tip: lowercase p = paste after/below, capital P = paste before/above. " +
					"This same after-vs-before pattern (a/A, i/I, p/P) repeats all over Vim.",
				NewKeys:     []string{"P"},
				Start:       []string{"a", "b", "target", "c"},
				CursorStart: Pos{2, 0},
				Kind:        KindEdit,
				Target:      []string{"a", "b", "target", "target", "target", "c"},
				Par:         4,
			},
			{
				Title: "Relocate The Line",
				Instructions: "\"four\" is stuck in the wrong spot near the top — combine dd to " +
					"lift it out, G to jump to the very last line, and p to drop it back in at " +
					"the end where it belongs.",
				Tip: "Tip: dd doesn't just delete — it stores the line exactly like yy would, " +
					"ready for p or P whenever you need it.",
				Start:       []string{"one", "four", "two", "three"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"one", "two", "three", "four"},
				Par:         4,
			},
			{
				Title: "Double Swap",
				Instructions: "This list has two separate pairs of lines swapped out of order. " +
					"Find both mistakes and fix each one using dd and p to lift a line out and " +
					"drop it back in the right place.",
				Tip: "Tip: fix them one at a time, top to bottom — a dd/p pair right after each " +
					"other is usually enough to swap two neighboring lines.",
				Start:       []string{"apple", "cherry", "banana", "date", "fig", "honey", "grape"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"apple", "banana", "cherry", "date", "fig", "grape", "honey"},
				Par:         9,
			},
		},
	},
	{
		Number: 6,
		Week:   "Week 2: More Operators",
		Title:  "dw, cw, ciw, diw — Operators Meet Word Motions",
		Summary: "Operators (d, c, y) and motions (w) combine freely — you're not limited " +
			"to the dd/yy pairs from Day 5. dw deletes a word forward. cw changes one, but " +
			"has a real Vim quirk: it stops at the end of the current word instead of also " +
			"eating the space after it. ciw and diw sidestep that entirely by targeting the " +
			"whole word around the cursor, wherever inside it you happen to be.",
		Challenges: []Challenge{
			{
				Title: "Delete The Extra Word",
				Instructions: "You already know d deletes and w moves by word — combine them and " +
					"dw deletes from the cursor to the start of the next word in a single motion. " +
					"Two stray words, \"very\" and \"extremely\", need to go from the line below.",
				Tip: "Tip: like dd, this combo takes a count too — d3w deletes three words " +
					"forward in one shot.",
				NewKeys:     []string{"dw"},
				Start:       []string{"quick very extremely fast car today"},
				CursorStart: Pos{0, 6},
				Kind:        KindEdit,
				Target:      []string{"quick fast car today"},
				Par:         4,
			},
			{
				Title: "The cw Quirk",
				Instructions: "c changes text just like d deletes it, but cw has a special " +
					"exception: it only replaces up to the end of the word under the cursor — it " +
					"leaves the trailing space alone, unlike dw. Fix both typos below, \"quikc\" " +
					"and \"borwn\", using cw.",
				Tip: "Tip: this exception exists so cw feels natural for fixing typos — without " +
					"it, retyping a word would also eat the space after it, forcing you to add it " +
					"back yourself.",
				NewKeys:     []string{"cw"},
				Start:       []string{"quikc borwn fox"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"quick brown fox"},
				Par:         17,
			},
			{
				Title: "Change Inner Word",
				Instructions: "ciw changes the whole word your cursor is inside, no matter which " +
					"letter you're on — unlike cw, you don't need to be sitting on the word's " +
					"first letter. Fix both \"woord\" and \"tiem\" below using ciw, starting from " +
					"the middle of each.",
				Tip: "Tip: iw also works with d and y (diw, yiw) — because it grabs just the " +
					"word itself, it's more predictable than w-based motions when the cursor " +
					"isn't sitting on a word's edge.",
				NewKeys:     []string{"ciw"},
				Start:       []string{"fix the woord and tiem now"},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target:      []string{"fix the word and time now"},
				Par:         18,
			},
			{
				Title: "The diw Gap",
				Instructions: "diw deletes just the inner word, same target as ciw but without " +
					"dropping into Insert mode afterward — and it leaves the surrounding spaces " +
					"exactly where they were, unlike daw which you'll meet tomorrow. Delete the " +
					"second \"there\" below with diw and watch the double space it leaves behind.",
				Tip: "Tip: that leftover double space is exactly the problem daw was built to " +
					"solve — keep this one in mind for Day 7.",
				NewKeys:     []string{"diw"},
				Start:       []string{"hello there there friend"},
				CursorStart: Pos{0, 14},
				Kind:        KindEdit,
				Target:      []string{"hello there  friend"},
				Par:         3,
			},
			{
				Title: "Word Surgery",
				Instructions: "Three lines, three different problems: an extra word to strip out, " +
					"a typo sitting at a word's first letter, and a typo buried in the middle of " +
					"a word. Pick the right tool for each — dw, cw, or ciw.",
				Tip: "Tip: if the cursor is already on the word's first letter, cw is enough — " +
					"only reach for ciw when you're starting from somewhere in the middle.",
				Start: []string{
					"remove very this word here",
					"please fix kwick now",
					"clean woord please",
				},
				CursorStart: Pos{0, 7},
				Kind:        KindEdit,
				Target: []string{
					"remove this word here",
					"please fix quick now",
					"clean word please",
				},
				Par: 22,
			},
		},
	},
	{
		Number: 7,
		Week:   "Week 2: More Operators",
		Title:  "daw, D, C — Around Word & End-of-Line Shortcuts",
		Summary: "iw grabs just the word itself; aw grabs the word plus the whitespace " +
			"around it, so daw cleans up after itself without leaving a double space behind " +
			"like diw did yesterday. D and C are shortcuts for d$ and c$ — delete or change " +
			"everything from the cursor to the end of the line in a single keystroke.",
		Challenges: []Challenge{
			{
				Title: "Delete Around Word",
				Instructions: "aw is like iw but also grabs the whitespace touching the word, so " +
					"daw removes a whole word cleanly without leaving a stray double space behind " +
					"like diw would. Use daw twice to remove both \"cat\" and \"junk\" from the " +
					"line below.",
				Tip: "Tip: yaw and caw follow the same pattern — y/c/d all pair with either the " +
					"i (inner) or a (around) text objects.",
				NewKeys:     []string{"daw"},
				Start:       []string{"please cat dog junk run fast"},
				CursorStart: Pos{0, 7},
				Kind:        KindEdit,
				Target:      []string{"please dog run fast"},
				Par:         7,
			},
			{
				Title: "Delete To End Of Line",
				Instructions: "D deletes from the cursor all the way to the end of the line — a " +
					"shortcut for d$. Both lines below have junk tacked onto the end starting at " +
					"the same column — use D on each to strip it off.",
				Tip: "Tip: D is one of a handful of capital-letter shortcuts for a lowercase " +
					"operator plus $ — C works the same way for c$, as you'll see next.",
				NewKeys:     []string{"D"},
				Start:       []string{"good stuff###delete me here", "nice today###forget this bit"},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target:      []string{"good stuff", "nice today"},
				Par:         3,
			},
			{
				Title: "Change To End Of Line",
				Instructions: "C deletes from the cursor to the end of the line and drops you into " +
					"Insert mode there — a shortcut for c$. Use C to fix the typo \"yuo\" on the " +
					"first line, then fix \"konw\" on the second by retyping its whole tail.",
				Tip: "Tip: same after-vs-before family as D/C — think of them as \"the rest of " +
					"the line, gone\" plus whether you land in Insert mode afterward.",
				NewKeys:     []string{"C"},
				Start:       []string{"nice to meet yuo", "great to konw you"},
				CursorStart: Pos{0, 13},
				Kind:        KindEdit,
				Target:      []string{"nice to meet you", "great to know you"},
				Par:         20,
			},
			{
				Title: "Combo Cleanup",
				Instructions: "One line, two fixes: a stray word to remove with daw, then junk " +
					"trailing after \"###\" to strip off with D. Use word motions to travel " +
					"between the two.",
				Tip: "Tip: \"###\" counts as its own word since it's a run of punctuation — w " +
					"lands you right at the start of it, exactly where D needs to strike.",
				Start:       []string{"please extra word here###junk tail data"},
				CursorStart: Pos{0, 7},
				Kind:        KindEdit,
				Target:      []string{"please word here"},
				Par:         6,
			},
			{
				Title: "Capstone",
				Instructions: "Two lines, both needing a different Day 7 tool: an extra word to " +
					"strip with daw on the first, then junk to trim with D on the second. Get " +
					"from one fix to the other using a word motion, not hjkl.",
				Tip: "Tip: b from inside a run of punctuation snaps to the start of that same " +
					"token — handy for lining the cursor up right before a \"###\" marker.",
				Start:       []string{"good words extra stuff here", "trim this###junk data now"},
				CursorStart: Pos{0, 11},
				Kind:        KindEdit,
				Target:      []string{"good words stuff here", "trim this"},
				Par:         6,
			},
		},
	},
	{
		Number: 8,
		Week:   "Week 2: Checkpoint",
		Title:  "Boss Challenge — Everything From Days 1-7",
		Summary: "No new keys today. Five tasks, each solvable only by combining motions, " +
			"operators, and Insert-mode commands from Days 1-7 — no hand-holding, no single " +
			"correct sequence spelled out for you.",
		Challenges: []Challenge{
			{
				Title: "Fix Two Typos",
				Instructions: "Two lines below have typos. Move between lines and fix each word " +
					"— figure out which operator and motion combo gets the job done fastest.",
				Tip: "Tip: you don't have to delete and retype an entire word to fix a typo " +
					"inside it — Day 6 covered exactly this.",
				Start: []string{
					"quikc brown fox",
					"jumps ovver the",
					"lazy dog now",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"quick brown fox",
					"jumps over the",
					"lazy dog now",
				},
				Par: 17,
			},
			{
				Title: "Sort Them Out",
				Instructions: "These three lines are out of alphabetical order. Move the odd one " +
					"out into its correct place — no retyping, just cut and paste.",
				Tip: "Tip: think back to Day 5's dd/p/P — deleting a line doesn't just remove " +
					"it, it stores it for pasting elsewhere.",
				Start:       []string{"banana", "apple", "cherry"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"apple", "banana", "cherry"},
				Par:         5,
			},
			{
				Title: "Finish The Story",
				Instructions: "The first line is missing a word at the end, and the story is " +
					"missing its last line entirely. Fix both using Insert-mode commands from " +
					"Day 3.",
				Tip: "Tip: A jumps to end-of-line before inserting, o opens a fresh line below " +
					"— neither needs any cursor-walking first.",
				Start:       []string{"Hello", "Goodbye friend"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"Hello world", "Goodbye friend", "See you soon"},
				Par:         23,
			},
			{
				Title: "One Line, Two Fixes",
				Instructions: "Both typos live on this single line. Fix the first one, then get " +
					"to the second using a word motion instead of hjkl — then fix that one too.",
				Tip: "Tip: w still works right where Normal mode left the cursor after Insert " +
					"mode ends — no need to reposition first.",
				Start:       []string{"The quikc brown fox jumps ovver lazy"},
				CursorStart: Pos{0, 4},
				Kind:        KindEdit,
				Target:      []string{"The quick brown fox jumps over lazy"},
				Par:         19,
			},
			{
				Title: "Final Boss",
				Instructions: "One line is in the wrong place, one word has a typo, one word " +
					"shouldn't be there, and one line has junk tacked onto the end. Fix all four " +
					"problems using whatever combination of Day 1-7 tools gets there fastest.",
				Tip: "Tip: reorder first with dd/p, then work top to bottom fixing what's left " +
					"— it's easier to fix typos in a list that's already in the right order.",
				Start: []string{
					"banana",
					"apple",
					"quikc brown fox",
					"jumps extra over",
					"the lazy dog###done",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"apple",
					"banana",
					"quick brown fox",
					"jumps over",
					"the lazy dog",
				},
				Par: 21,
			},
		},
	},
	{
		Number: 9,
		Week:   "Week 2: More Operators",
		Title:  "r and ~ — Replace One Character & Toggle Case",
		Summary: "r{char} replaces the single character under the cursor with whatever you " +
			"type next, without ever entering Insert mode — the fastest fix for a one-letter " +
			"typo. Prefix it with a count, like 5rX, and it replaces that many characters in " +
			"one shot. ~ toggles the case of the character under the cursor and steps forward " +
			"one column, so mashing it repeatedly walks case-toggles across a whole run of " +
			"letters.",
		Challenges: []Challenge{
			{
				Title: "Replace One Letter",
				Instructions: "r followed by a single character swaps in that character for " +
					"whatever's under the cursor and leaves you in Normal mode — no i, no esc. " +
					"\"fovorite\" below has one wrong letter; fix it with a single r.",
				Tip: "Tip: r is the fastest way to fix a one-letter typo — x then i then esc " +
					"does the same job in three separate steps.",
				NewKeys:     []string{"r"},
				Start:       []string{"This is my fovorite meal"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"This is my favorite meal"},
				Par:         6,
			},
			{
				Title: "Redact With A Count",
				Instructions: "A count in front of r replaces that many characters in one move — " +
					"5rX replaces the next 5 characters under and after the cursor with X, all at " +
					"once. Use it to redact the 5-digit number below into XXXXX.",
				Tip: "Tip: same count trick as 3dd or 7gg — r just applies it to how many " +
					"characters get replaced instead of how many lines get affected.",
				Start:       []string{"the code is 99999 today"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"the code is XXXXX today"},
				Par:         6,
			},
			{
				Title: "Toggle One Letter's Case",
				Instructions: "~ flips the case of the character under the cursor — upper " +
					"becomes lower, lower becomes upper — and moves the cursor forward one " +
					"column afterward. \"worLd\" below has one letter in the wrong case; fix it " +
					"with a single ~.",
				Tip: "Tip: unlike r, ~ never asks for a second keystroke — it always toggles " +
					"whatever's under the cursor right away.",
				NewKeys:     []string{"~"},
				Start:       []string{"Hello worLd"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"Hello world"},
				Par:         5,
			},
			{
				Title: "Toggle A Whole Run",
				Instructions: "Because ~ steps forward after every toggle, you can fix a whole " +
					"run of wrongly-cased letters just by mashing it — no need to move the " +
					"cursor between each one. \"HELLO\" below is shouting; walk ~ across all 5 " +
					"letters to calm it down.",
				Tip: "Tip: this same step-forward-after-each-press pattern is what let x clear " +
					"a run of characters back on Day 4 — ~ just toggles case instead of deleting.",
				Start:       []string{"please say HELLO now"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"please say hello now"},
				Par:         7,
			},
			{
				Title: "Capstone",
				Instructions: "Three lines, three different jobs: two adjacent wrong letters to " +
					"fix with r on the first line, a shouted word to calm down with ~ on the " +
					"second, and a number to redact with a count-prefixed r on the third.",
				Tip: "Tip: b snaps back to the start of the word your cursor is already inside " +
					"— handy for lining up before mashing ~ across an all-caps word.",
				Start: []string{
					"we had figt yesterday",
					"keep it QUIET okay",
					"id number 88888 here",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"we had five yesterday",
					"keep it quiet okay",
					"id number XXXXX here",
				},
				Par: 21,
			},
		},
	},
	{
		Number: 10,
		Week:   "Week 2: More Operators",
		Title:  "s and S — Substitute Character & Substitute Line",
		Summary: "s deletes the character(s) under the cursor and drops you straight into " +
			"Insert mode — like x followed by i in one command, except unlike r you can type " +
			"back as many characters as you want instead of exactly one. A count in front, " +
			"like 3s, deletes that many characters before you start typing. S goes bigger: it " +
			"clears the entire line's content and drops you into Insert mode to type a full " +
			"replacement — a count in front, like 2S, merges that many lines into whatever " +
			"you type.",
		Challenges: []Challenge{
			{
				Title: "Swap And Type",
				Instructions: "s deletes the character under the cursor and puts you in Insert " +
					"mode so you can type a replacement of any length — not just the single " +
					"character r gives you. The line below has a digit that should be spelled " +
					"out; land on the \"5\" and use s to swap it for \"five\".",
				Tip: "Tip: s is really x then i as a single keystroke — but since it drops you " +
					"into Insert mode, you're free to type more (or fewer) characters than you " +
					"deleted.",
				NewKeys:     []string{"s"},
				Start:       []string{"I have 5 dogs"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"I have five dogs"},
				Par:         8,
			},
			{
				Title: "Count Before You Type",
				Instructions: "A count in front of s deletes that many characters before " +
					"dropping you into Insert mode — 3s deletes 3 characters, then whatever you " +
					"type replaces them. The number below is wrong; delete all three digits with " +
					"a count and type the fix.",
				Tip: "Tip: same count pattern as 5rX from Day 9 — the count says how many " +
					"characters to remove, except now you type the replacement yourself instead " +
					"of it being one repeated character.",
				Start:       []string{"The CPU is 100% busy"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"The CPU is 50% busy"},
				Par:         8,
			},
			{
				Title: "Rewrite The Line",
				Instructions: "S clears the whole line and drops you into Insert mode to type a " +
					"full replacement, no matter where on the line your cursor started. The line " +
					"below is a mess from end to end; use S to throw it all away and type a " +
					"clean one.",
				Tip: "Tip: S is a shortcut for typing cc (Vim's change-the-whole-line command) " +
					"— same result, one less keystroke, and you don't need to be at column 0 " +
					"first.",
				NewKeys:     []string{"S"},
				Start:       []string{"this line is a mess and wrong"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"this line is correct now"},
				Par:         26,
			},
			{
				Title: "Merge Two Into One",
				Instructions: "A count in front of S works just like it does for dd or yy — 2S " +
					"clears 2 lines and drops you into Insert mode once, so whatever you type " +
					"becomes a single line replacing both. Merge the two rough draft lines below " +
					"into the one clean line shown as the target.",
				Tip: "Tip: this is how you combine two lines into one without ever touching dd " +
					"— delete-and-retype in one motion instead of two separate steps.",
				Start:       []string{"draft line one", "draft line two"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"final combined line"},
				Par:         22,
			},
			{
				Title: "Capstone",
				Instructions: "Three lines, three jobs: a number to spell out with a " +
					"count-prefixed s, a whole line to rewrite with S, and a typo'd word to " +
					"replace with a count-prefixed s.",
				Tip: "Tip: S doesn't care what column you're on when you press it — land " +
					"anywhere on the line and it clears the whole thing.",
				Start: []string{
					"the score is 42 points",
					"this whole line is nonsense and needs replacing",
					"fix typo1 in here",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"the score is one hundred points",
					"clean sentence here",
					"fix this in here",
				},
				Par: 50,
			},
		},
	},
	{
		Number: 11,
		Week:   "Week 2: Checkpoint",
		Title:  "Boss",
		Summary: "No new keys today. Five tasks, each solvable only by combining motions, " +
			"operators, and Insert-mode commands from Days 1-10 — no hand-holding, no single " +
			"correct sequence spelled out for you.",
		Challenges: []Challenge{
			{
				Title: "Typo Triage",
				Instructions: "Three different problems on three lines: a single wrong letter, " +
					"a shouted word, and a digit that should be spelled out. Use r, ~, or s — " +
					"whichever fits each job.",
				Tip: "Tip: r only swaps in exactly one character — reach for s instead when " +
					"the fix needs a different number of characters than what's there now.",
				Start: []string{
					"the doy is big",
					"STOP yelling now",
					"buy 2 apples",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"the dog is big",
					"stop yelling now",
					"buy two apples",
				},
				Par: 19,
			},
			{
				Title: "Reorder And Merge",
				Instructions: "The first two lines are out of alphabetical order — fix that " +
					"without retyping anything. The last two lines are rough drafts that belong " +
					"together — merge them into the single clean line shown as the target.",
				Tip: "Tip: dd/p for the reorder, then a count in front of S for the merge — two " +
					"completely different tools for two completely different jobs.",
				Start: []string{
					"banana",
					"apple",
					"draft one",
					"draft two",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"apple",
					"banana",
					"final merged line",
				},
				Par: 24,
			},
			{
				Title: "Redact And Rewrite",
				Instructions: "The first line has a number that needs redacting, count and all. " +
					"The second line is garbage from start to finish and needs a full rewrite.",
				Tip: "Tip: a count in front of r redacts characters in place; S throws the " +
					"whole line away and starts fresh — pick whichever matches how much of the " +
					"line is actually wrong.",
				Start: []string{
					"ssn is 55555 today",
					"this whole line is garbage and should go",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"ssn is XXXXX today",
					"clean replacement line",
				},
				Par: 31,
			},
			{
				Title: "One Line, Many Fixes",
				Instructions: "This single line has three unrelated problems stacked on top of " +
					"each other: a shouted word, a one-letter typo, and a digit that should be " +
					"spelled out. Fix all three left to right.",
				Tip: "Tip: mash ~ across the shouted word first, then keep moving right with w " +
					"— no need to backtrack to the start of the line between fixes.",
				Start:       []string{"the CAT sit on 4 legs"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"the cat sat on four legs"},
				Par:         16,
			},
			{
				Title: "Final Boss",
				Instructions: "Two lines are in the wrong order, one line has a number to " +
					"redact, one line is shouting, and one line has junk tacked onto the end. " +
					"Fix all four problems using whatever combination of Day 1-10 tools gets " +
					"there fastest.",
				Tip: "Tip: reorder first with dd/p so the rest of the fixes read top to bottom " +
					"in a sensible order, then work your way down the list.",
				Start: []string{
					"cherry",
					"apple",
					"your pin is 12345 ok",
					"STOP shouting please",
					"the lazy dog###done",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"apple",
					"cherry",
					"your pin is XXXXX ok",
					"stop shouting please",
					"the lazy dog",
				},
				Par: 23,
			},
		},
	},
	{
		Number: 12,
		Week:   "Week 2: More Operators",
		Title:  "X and u — Delete Backward & Undo",
		Summary: "X deletes the character(s) immediately to the left of the cursor — the " +
			"mirror image of x, useful when you've landed one step past what actually needs " +
			"to go. A count in front, like 4X, deletes that many characters going left in one " +
			"shot. u undoes your most recent change and restores the buffer, and your cursor, " +
			"to exactly how it was before — press it again to keep walking back through " +
			"earlier edits, one step at a time.",
		Challenges: []Challenge{
			{
				Title: "Delete Backwards",
				Instructions: "X deletes the character immediately to the left of the cursor " +
					"and shifts everything after it left to fill the gap — it's x's mirror " +
					"image, looking backward instead of forward. The word below has one letter " +
					"typed twice; land just past the extra letter and press X to remove it.",
				Tip: "Tip: x deletes forward from the cursor, X deletes backward from the " +
					"cursor — same idea, opposite direction.",
				NewKeys:     []string{"X"},
				Start:       []string{"committ now"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"commit now"},
				Par:         2,
			},
			{
				Title: "Count Before X",
				Instructions: "A count in front of X deletes that many characters to the " +
					"left in one go, just like a count in front of x does going forward. The " +
					"filename below picked up a run of accidental leading zeros; move past " +
					"them and clear all of them with a single count-prefixed X.",
				Tip: "Tip: line up the count with exactly how many characters need to go — " +
					"5X removes 5, no more, no less.",
				Start:       []string{"file0000042.txt"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"file42.txt"},
				Par:         4,
			},
			{
				Title: "Undo Your Last Move",
				Instructions: "u undoes your most recent change and puts the buffer, and " +
					"your cursor, back exactly where they were before it happened. Try it: " +
					"delete \"the \" first with dw, see the wrong result, then press u to bring " +
					"it right back before removing the actual extra word, \"oops \", the same way.",
				Tip: "Tip: u doesn't just restore the text — it restores your cursor position " +
					"too, so you land right back where the mistaken edit started.",
				NewKeys:     []string{"u"},
				Start:       []string{"oops the plan is solid"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"the plan is solid"},
				Par:         7,
			},
			{
				Title: "Walk Back Two Steps",
				Instructions: "u isn't a one-shot rewind — press it again and it undoes the " +
					"change before that one, walking you back through your history one step " +
					"at a time. Delete all three words below one at a time with dw, then press " +
					"u twice to walk back exactly two of those deletions.",
				Tip: "Tip: each u undoes exactly one previous change, however small — three " +
					"edits forward means you can walk back up to three steps.",
				Start:       []string{"one two three four"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"two three four"},
				Par:         8,
			},
			{
				Title: "Capstone",
				Instructions: "Three lines, three tools: a doubled letter to fix with X, a " +
					"run of leading zeros to clear with a count-prefixed X, and one more spot " +
					"to practice the make-a-mistake-then-u trick from before.",
				Tip: "Tip: on the last line, delete the wrong word first on purpose so u has " +
					"something real to undo — then finish with the correct one.",
				Start: []string{
					"appple is red",
					"000007 items left",
					"oops the goal succeeds",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"apple is red",
					"7 items left",
					"the goal succeeds",
				},
				Par: 18,
			},
		},
	},
	{
		Number: 13,
		Week:   "Week 2: More Operators",
		Title:  "0, ^, and $ — Line Start, First Non-Blank, and Line End",
		Summary: "Three motions for jumping along a single line. 0 always jumps to the " +
			"true first column, even if that's blank indentation. ^ jumps to the first " +
			"non-blank character instead, skipping any leading spaces or tabs — the column " +
			"where the line's real content begins. $ jumps to the last character of the " +
			"line, and a count in front, like 3$, jumps to the end of the line that many " +
			"rows down. All three combine with d, c, and y exactly like any other motion — " +
			"d$ deletes to the end of the line (it's what D is secretly shorthand for), and " +
			"c^ deletes back to the first real character and drops you into Insert mode to " +
			"type a replacement.",
		Challenges: []Challenge{
			{
				Title: "The True Beginning",
				Instructions: "0 always jumps to column zero — the absolute start of the " +
					"line — even when the line begins with blank space. Land your cursor on " +
					"the very first character of the line below, blank or not.",
				Tip: "Tip: 0 doesn't care what's actually on the line. It just goes to the " +
					"first column, every time.",
				NewKeys:     []string{"0"},
				Start:       []string{"    indented text here"},
				CursorStart: Pos{0, 15},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 0},
				Par:         1,
			},
			{
				Title: "Skip The Blank Space",
				Instructions: "^ jumps to the first non-blank character on the line — the " +
					"spot where the real content starts, ignoring any leading indentation. " +
					"Land on the \"t\" that begins the actual sentence below.",
				Tip: "Tip: ^ is what you want when a line is indented and you don't care " +
					"about the blank space in front of it — 0 would leave you sitting on a " +
					"space instead.",
				NewKeys:     []string{"^"},
				Start:       []string{"    the real text starts here"},
				CursorStart: Pos{0, 28},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 4},
				Par:         1,
			},
			{
				Title: "Straight To The Last Line",
				Instructions: "$ jumps to the last character of the line. A count in front " +
					"works like it does for G, but line-local: 3$ jumps straight to the end " +
					"of the 3rd line down instead of the current one. Land on the very last " +
					"character of the 3rd line below in one shot.",
				Tip: "Tip: 3$ means \"the end of 3 lines down,\" not \"3 characters to the " +
					"right of the end\" — the count picks which line, not a distance along it.",
				NewKeys: []string{"$"},
				Start: []string{
					"short",
					"medium line here",
					"the longest line of them all",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{2, 27},
				Par:         2,
			},
			{
				Title: "Cut Both Ends",
				Instructions: "d$ deletes from the cursor to the end of the line — the same " +
					"result as D, just spelled out as an explicit motion. c^ deletes back to " +
					"the first non-blank character and drops you into Insert mode. The first " +
					"line below has trailing junk to cut with d$; the second has a garbage " +
					"prefix (after its indentation) to replace with c^.",
				Tip: "Tip: D really is just d$ under the hood — now you've typed both and can " +
					"see they do the exact same thing.",
				Start: []string{
					"keep this part but not this trailing junk",
					"    XXXX real content here",
				},
				CursorStart: Pos{0, 14},
				Kind:        KindEdit,
				Target: []string{
					"keep this part",
					"    fix: real content here",
				},
				Par: 13,
			},
			{
				Title: "Capstone",
				Instructions: "Three lines, three jobs: prepend a note at the absolute start " +
					"of an indented line with 0, cut trailing junk from the second line with " +
					"d$, and replace a garbage prefix on the third line with c^.",
				Tip: "Tip: 0 lands before the indentation, ^ lands after it — pick whichever " +
					"one the job actually needs.",
				Start: []string{
					"    fix the indentation on this line",
					"keep this good part TRASH from here onward",
					"    XXXX apply the fix here too",
				},
				CursorStart: Pos{0, 20},
				Kind:        KindEdit,
				Target: []string{
					"NOTE:     fix the indentation on this line",
					"keep this good part ",
					"    DONE: apply the fix here too",
				},
				Par: 29,
			},
		},
	},
	{
		Number: 14,
		Week:   "Week 2: Final Boss",
		Title:  "Final Boss — Everything From Days 1-13",
		Summary: "No new keys today — this is the big one. Every challenge below leans on " +
			"tools from across the last two weeks: motions, insert-mode commands, the delete " +
			"and yank operators, word and around-word text objects, replace and case-toggle, " +
			"substitute, backward-delete and undo, and the line-start/first-non-blank/line-end " +
			"motions from yesterday. The last challenge is a genuinely long, multi-part " +
			"document with a dozen separate problems to fix — expect it to take real time, not " +
			"a quick pass.",
		Challenges: []Challenge{
			{
				Title: "Warm It Up",
				Instructions: "Three small jobs to shake off the rust: fix the last word of " +
					"the first line with cw, delete the second line entirely with dd, and " +
					"duplicate the third line below itself with yy and p.",
				Tip: "Tip: nothing here is new — it's the exact same dd and yy/p from Day 5, " +
					"and the same cw from Day 6.",
				Start: []string{
					"jump to the correct spot",
					"delete this whole line please",
					"keep me",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"jump to the correct target",
					"keep me",
					"keep me",
				},
				Par: 17,
			},
			{
				Title: "Bigger Fixes",
				Instructions: "Five lines, five different tools: mash ~ to fix a shouted " +
					"word, use ciw to fix a misspelling, use d$ to trim a trailing phrase, " +
					"use s to replace a number with a spelled-out word, and use a " +
					"count-prefixed r to redact a number in place.",
				Tip: "Tip: match the tool to the job — ciw for a whole misspelled word, s " +
					"when the replacement is a different length than the original, r when " +
					"it's a straight character-for-character swap.",
				Start: []string{
					"STOP yelling immediately now",
					"fix this speling mistake here",
					"cut the ending right here needless extra words",
					"5 apples were bought today",
					"buy 100 items",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"stop yelling immediately now",
					"fix this spelling mistake here",
					"cut the ending right here",
					"three apples were bought today",
					"buy XXX items",
				},
				Par: 45,
			},
			{
				Title: "Line Surgery",
				Instructions: "The three lines below are in the wrong order — fix that with " +
					"dd and p, no retyping. Then add a fourth line at the very end with G and o.",
				Tip: "Tip: this is the same dd/p reordering trick from the Week 2 checkpoints, " +
					"just with fewer lines to juggle.",
				Start: []string{
					"third",
					"first",
					"second",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"first",
					"second",
					"third",
					"fourth",
				},
				Par: 13,
			},
			{
				Title: "Indentation Cleanup",
				Instructions: "Three indented lines, three fixes from yesterday: replace a " +
					"garbage prefix with c^, trim a trailing phrase with d$, and prepend a " +
					"marker at the true start of a line with 0 and i.",
				Tip: "Tip: c^ and d$ don't touch the leading indentation itself — only 0 " +
					"reaches all the way back to column zero.",
				Start: []string{
					"    XXXX keep this part clean",
					"keep this text but###trim the rest",
					"    old note more content stays",
				},
				CursorStart: Pos{0, 20},
				Kind:        KindEdit,
				Target: []string{
					"    TODO: keep this part clean",
					"keep this text but",
					"*     old note more content stays",
				},
				Par: 25,
			},
			{
				Title: "Final Boss",
				Instructions: "A whole messy document, twelve lines, one pass: a garbage " +
					"prefix to replace, an accidentally duplicated line to remove, a typo to " +
					"fix, an entire line of junk to delete, a number to redact, a shouted " +
					"line to fix, a trailing phrase to trim, three lines to fully reorder, " +
					"and a brand new line to add at the end. Nothing here is new — it's every " +
					"tool from Days 1-13, back to back, on a document big enough to actually " +
					"take some time to fix properly.",
				Tip: "Tip: work top to bottom, one problem at a time, and don't rush — this " +
					"one's meant to take a few minutes, not a few seconds.",
				Start: []string{
					"    XXXX this document has a messy header",
					"TWO lines got duplicated by accident",
					"TWO lines got duplicated by accident",
					"fix the tyop in this sentence",
					"delete this entire line of garbage text",
					"the count is 999 and needs redacting",
					"STOP SHOUTING in the summary line",
					"extra junk needless trailing words here",
					"cherry",
					"banana",
					"apple",
					"insert a new closing line below",
				},
				CursorStart: Pos{0, 20},
				Kind:        KindEdit,
				Target: []string{
					"    TODO: this document has a messy header",
					"TWO lines got duplicated by accident",
					"fix the typo in this sentence",
					"the count is XXX and needs redacting",
					"stop shouting in the summary line",
					"extra junk",
					"apple",
					"banana",
					"cherry",
					"insert a new closing line below",
					"closing line added at the end",
				},
				Par: 100,
			},
		},
	},
	{
		Number: 15,
		Week:   "Week 3: Counts & Text Objects",
		Title:  "Counts — Repeat Any Motion or Operator in One Shot",
		Summary: "You've already used a count in front of r, s, S, and X — today it's " +
			"official: almost any motion or operator can take a number in front of it to " +
			"repeat that action that many times in a single command. 4w jumps forward four " +
			"words, 3dd deletes three whole lines, 3yy copies three lines at once. A count " +
			"can also sit between an operator and its motion — d3w and 3dw both delete " +
			"exactly three words, because the two counts multiply together no matter which " +
			"side of the operator they're on.",
		Challenges: []Challenge{
			{
				Title: "Counted Word Jumps",
				Instructions: "A number before w repeats it that many times in one command — " +
					"4w is the same as w w w w, just two keystrokes instead of four. It works " +
					"backward too: 2b jumps back two words. Use 4w to jump forward four words, " +
					"then 2b to jump back two, and land exactly on the target word.",
				Tip: "Tip: think of the count as a multiplier you attach to a motion you " +
					"already know, not a new command to learn.",
				Start:       []string{"alpha beta gamma delta epsilon zeta eta theta"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 11},
				Par:         4,
			},
			{
				Title: "Delete Three Lines",
				Instructions: "A count before dd deletes that many whole lines starting from " +
					"the cursor, in one command instead of repeating dd over and over. The list " +
					"below has three junk lines on top; clear all three with a single 3dd.",
				Tip: "Tip: 3dd is dd's version of the count trick — same idea as 3yy for " +
					"copying or 4X for deleting characters, just linewise.",
				NewKeys: []string{"3dd"},
				Start: []string{
					"remove me",
					"remove me too",
					"remove me three",
					"keep this line",
					"keep this line too",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"keep this line",
					"keep this line too",
				},
				Par: 3,
			},
			{
				Title: "Copy A Block Elsewhere",
				Instructions: "3yy copies three lines starting from the cursor into a single " +
					"block, just like 3dd deletes three — except it leaves the original lines " +
					"untouched. Yank the three template lines with 3yy, jump to the very bottom " +
					"with G, then paste the whole block after the last line with p.",
				Tip: "Tip: G moves you to the last line first so the pasted block lands " +
					"cleanly at the end, instead of splitting in the middle of the block you " +
					"just copied.",
				Start: []string{
					"template line one",
					"template line two",
					"template line three",
					"separator",
					"END",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"template line one",
					"template line two",
					"template line three",
					"separator",
					"END",
					"template line one",
					"template line two",
					"template line three",
				},
				Par: 5,
			},
			{
				Title: "Two Counts, Same Result",
				Instructions: "A count can sit on either side of an operator and it works the " +
					"same way: d3w and 3dw both delete exactly three words, because the counts " +
					"multiply together no matter which one comes first. Use d3w to clear three " +
					"words from the first line, then use 3dw to do the same thing the other way " +
					"on the second line.",
				Tip: "Tip: this also means d2d and 2dd are the same command — pick whichever " +
					"order reads more naturally to you in the moment.",
				NewKeys: []string{"d3w", "3dw"},
				Start: []string{
					"delete these three words please",
					"remove all four items now please",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"words please",
					"items now please",
				},
				Par: 8,
			},
			{
				Title: "Capstone",
				Instructions: "One document, three count tricks back to back: clear three " +
					"junk lines off the top with 3dd, strip three throwaway words off the next " +
					"line with d3w, then copy the three-line template block near the bottom with " +
					"3yy and paste it after the very last line with G and p.",
				Tip: "Tip: nothing here is new — it's the same three commands from this day's " +
					"earlier challenges, just chained together on one bigger document.",
				Start: []string{
					"junk line one to delete",
					"junk line two to delete",
					"junk line three to delete",
					"keep this line intact",
					"please remove three items now",
					"template block start",
					"template block middle",
					"template block end",
					"footer marker",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"keep this line intact",
					"items now",
					"template block start",
					"template block middle",
					"template block end",
					"footer marker",
					"template block start",
					"template block middle",
					"template block end",
				},
				Par: 15,
			},
		},
	},
	{
		Number: 16,
		Week:   "Week 3: Counts & Text Objects",
		Title:  "i( and i\" — Text Objects For Brackets & Quotes",
		Summary: "iw and aw grab a word, but text objects aren't limited to words: i( grabs " +
			"everything inside a pair of parentheses, a( grabs the parentheses too, and i\" " +
			"does the same trick for quoted text. Your cursor doesn't need to sit right next " +
			"to the ( or the \" either — anywhere inside the pair works, because Vim scans " +
			"outward to find the enclosing delimiters for you.",
		Challenges: []Challenge{
			{
				Title: "Clear The Call's Arguments",
				Instructions: "di( deletes everything between a pair of parentheses without " +
					"touching the parentheses themselves — same idea as diw, just a new pair " +
					"of delimiters instead of word boundaries. Your cursor is already sitting " +
					"inside the parens; press di( to empty them out.",
				Tip: "Tip: you don't need to be right after the ( — di( works from anywhere " +
					"inside the pair, because Vim searches outward for the nearest enclosing " +
					"( and ).",
				NewKeys:     []string{"di("},
				Start:       []string{"run(old broken code)"},
				CursorStart: Pos{0, 5},
				Kind:        KindEdit,
				Target:      []string{"run()"},
				Par:         3,
			},
			{
				Title: "Remove The Whole Aside",
				Instructions: "a( means 'around parentheses' — pair it with d and it deletes " +
					"the parentheses themselves along with everything inside them, unlike i( " +
					"which leaves the parens standing. Use da( to cut the whole aside out of " +
					"this line.",
				Tip: "Tip: i( vs a( is the same inside/around split as iw vs aw — i( leaves " +
					"the delimiters behind, a( takes them with it.",
				NewKeys:     []string{"da("},
				Start:       []string{"keep this (delete this aside) part"},
				CursorStart: Pos{0, 20},
				Kind:        KindEdit,
				Target:      []string{"keep this  part"},
				Par:         3,
			},
			{
				Title: "Swap The Argument",
				Instructions: "c works with parenthesis text objects too: ci( clears " +
					"everything inside the parens and drops you into Insert mode right where " +
					"they used to be, ready to type a replacement.",
				Tip: "Tip: ci( works exactly like ciw — the c, i pattern never changes, only " +
					"the delimiter after it does.",
				NewKeys:     []string{"ci("},
				Start:       []string{"print(placeholder)"},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target:      []string{"print(42)"},
				Par:         6,
			},
			{
				Title: "Rename The Quoted File",
				Instructions: "Quotes work as a text object too, with the exact same i/a " +
					"rule: ci\" clears whatever sits between a pair of double quotes and " +
					"drops you into Insert mode to type a replacement, just like ci( did for " +
					"parentheses.",
				Tip: "Tip: i\" and a\" follow the identical inside/around split as every " +
					"other text object you've learned — only the delimiter changes.",
				NewKeys:     []string{"ci\""},
				Start:       []string{"open the file \"notes.txt\" and read it"},
				CursorStart: Pos{0, 18},
				Kind:        KindEdit,
				Target:      []string{"open the file \"config.yaml\" and read it"},
				Par:         15,
			},
			{
				Title: "Capstone",
				Instructions: "Four lines, every bracket and quote trick from today: empty " +
					"a call's arguments with di(, cut a whole parenthetical aside with da(, " +
					"swap a call's argument with ci(, and rename a quoted value with ci\". " +
					"Use f( or f\" to jump onto the delimiter first on lines where your " +
					"cursor isn't already inside it.",
				Tip: "Tip: landing exactly on the ( or \" works just as well as landing " +
					"somewhere inside it — Vim only cares that you're inside the pair, " +
					"including its edges.",
				Start: []string{
					"run(old broken code)",
					"keep this (delete this aside) part",
					"print(placeholder)",
					"open the file \"notes.txt\" and read it",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"run()",
					"keep this  part",
					"print(42)",
					"open the file \"config.yaml\" and read it",
				},
				Par: 41,
			},
		},
	},
	{
		Number: 17,
		Week:   "Week 3: Counts & Text Objects",
		Title:  "i[ — Text Objects For Square Brackets",
		Summary: "Square brackets get the exact same text-object treatment as parentheses " +
			"and quotes: i[ grabs everything inside a pair of [ ], a[ grabs the brackets " +
			"too, and c/d work with either one just like they did for i(/a(.",
		Challenges: []Challenge{
			{
				Title: "Swap The Array Value",
				Instructions: "ci[ clears everything inside a pair of square brackets and " +
					"drops you into Insert mode right where they used to be — same c, i " +
					"pattern as ci(, just a new delimiter. Replace the value inside the " +
					"brackets with 42.",
				Tip: "Tip: di[/da[ delete inside/around square brackets the same way " +
					"di(/da( did for parentheses — only the bracket shape changes, the " +
					"i/a rule never does.",
				NewKeys:     []string{"ci["},
				Start:       []string{"data[old_value]"},
				CursorStart: Pos{0, 6},
				Kind:        KindEdit,
				Target:      []string{"data[42]"},
				Par:         6,
			},
		},
	},
	{
		Number: 18,
		Week:   "Week 3: Counts & Text Objects",
		Title:  "i{ and i' — Finishing The Text-Object Family",
		Summary: "Two delimiters left to round out the set: curly braces and single quotes " +
			"follow the exact same i/a rule as every bracket and quote you've already " +
			"learned. Square brackets also get their full di[/da[ treatment here, not just " +
			"ci[ from yesterday — after today you'll have every text-object delimiter this " +
			"engine supports: ( ) \" [ ] { } '.",
		Challenges: []Challenge{
			{
				Title: "Empty The Index",
				Instructions: "di[ deletes everything between a pair of square brackets " +
					"without touching the brackets themselves — the same di( trick from " +
					"Day 16, just a new delimiter. Your cursor is already inside the " +
					"brackets; press di[ to empty them out.",
				Tip: "Tip: yesterday you only used ci[ — di[ and da[ work exactly the " +
					"same way on square brackets as they already do on parentheses.",
				NewKeys:     []string{"di["},
				Start:       []string{"queue[stale_id]"},
				CursorStart: Pos{0, 8},
				Kind:        KindEdit,
				Target:      []string{"queue[]"},
				Par:         3,
			},
			{
				Title: "Remove The Whole Index",
				Instructions: "da[ takes the brackets with it, unlike di[ which leaves " +
					"them standing — the same inside/around split you already know from " +
					"da( and da\". Use da[ to cut the whole indexed aside out of this line.",
				Tip: "Tip: i vs a means the exact same thing no matter which delimiter " +
					"follows it — inside leaves the delimiters, around takes them too.",
				NewKeys:     []string{"da["},
				Start:       []string{"keep this [delete this index] part"},
				CursorStart: Pos{0, 15},
				Kind:        KindEdit,
				Target:      []string{"keep this  part"},
				Par:         3,
			},
			{
				Title: "Swap The Setting",
				Instructions: "Curly braces are a text object too: ci{ clears everything " +
					"inside a pair of { } and drops you into Insert mode right where they " +
					"used to be, ready to type a replacement — same c, i pattern as ci( " +
					"and ci[, just a new delimiter shape.",
				Tip: "Tip: di{ and da{ delete inside/around curly braces the same way " +
					"their ( and [ cousins do, if you ever need delete instead of change.",
				NewKeys:     []string{"ci{"},
				Start:       []string{"config{old_value}"},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target:      []string{"config{debug=true}"},
				Par:         14,
			},
			{
				Title: "Rename The Single-Quoted Value",
				Instructions: "Single quotes work exactly like double quotes did on Day " +
					"16: ci' clears whatever sits between a pair of single quotes and " +
					"drops you into Insert mode to type a replacement.",
				Tip: "Tip: di' and da' exist too, following the same rule as every other " +
					"delimiter — this is the last new text object in the family.",
				NewKeys:     []string{"ci'"},
				Start:       []string{"author = 'old_name'"},
				CursorStart: Pos{0, 13},
				Kind:        KindEdit,
				Target:      []string{"author = 'Devansh'"},
				Par:         11,
			},
			{
				Title: "Capstone",
				Instructions: "Six lines, all five text-object delimiters you now know: " +
					"clear a bracketed index with di[, cut a parenthetical aside with " +
					"da(, swap a call's argument with ci(, rename a quoted path with " +
					"ci\", rename a quoted name with ci', and swap a brace-wrapped " +
					"setting with ci{. Use f followed by the delimiter to jump onto it " +
					"first on every line, since your cursor starts each line at column 0.",
				Tip: "Tip: landing exactly on the opening delimiter works just as well " +
					"as landing somewhere inside it, for every one of these pairs.",
				Start: []string{
					"fetch[stale_cache]",
					"note (delete this aside) here",
					"print(placeholder)",
					"path = \"old/file.txt\"",
					"greet('old_name')",
					"user{role=guest}",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"fetch[]",
					"note  here",
					"print(42)",
					"path = \"new.txt\"",
					"greet('Devansh')",
					"user{role=admin}",
				},
				Par: 70,
			},
		},
	},
	{
		Number: 19,
		Week:   "Week 3: Counts & Text Objects",
		Title:  "yiw, yi(, yi\" — Yank Text Objects",
		Summary: "y takes a text object too, and it works exactly like d and c already do: " +
			"yiw copies a word, yi( copies what's inside a pair of parens, yi\" copies a " +
			"quoted string — none of it touches the buffer until you paste it with p or P. " +
			"The one catch: d and c both overwrite whatever y just stored, so if you need to " +
			"delete something else before you paste, delete it first and yank second.",
		Challenges: []Challenge{
			{
				Title: "Duplicate The Item",
				Instructions: "yiw copies a word into the register without changing anything " +
					"on screen — same iw text object as ciw and diw, but nothing gets deleted " +
					"this time. Yank \"milk\" with yiw, jump to the end of the line with $, " +
					"type \", \" to start a second copy, then press p to paste the word back in.",
				Tip: "Tip: y never touches the buffer by itself — it just fills the register, " +
					"so you're free to move anywhere on the line (or off it) before pasting " +
					"with p or P.",
				NewKeys:     []string{"yiw"},
				Start:       []string{"shopping list: milk"},
				CursorStart: Pos{0, 15},
				Kind:        KindEdit,
				Target:      []string{"shopping list: milk, milk"},
				Par:         9,
			},
			{
				Title: "Copy The Status Value",
				Instructions: "Two lines, same shape, different value. Delete the wrong word " +
					"on the second line first with diw, then jump up to the first line and " +
					"yank the correct one with yiw, then come back down and paste it into " +
					"the gap with p.",
				Tip: "Tip: order matters here — d and c both overwrite whatever y just put " +
					"in the register, so if you yank first and then delete something else, " +
					"your copy is gone before you can paste it. Delete first, yank second, " +
					"paste last.",
				Start:       []string{"status = active", "status = pending"},
				CursorStart: Pos{0, 9},
				Kind:        KindEdit,
				Target:      []string{"status = active", "status = active"},
				Par:         10,
			},
			{
				Title: "Copy The Scale Factor",
				Instructions: "Same idea with parentheses: clear the wrong value on the " +
					"second line first with di(, then go up and yank the right one with " +
					"yi(, then come back down to paste it in — but notice the paste key " +
					"changes here.",
				Tip: "Tip: di( leaves your cursor sitting on the closing ), so p (paste " +
					"after) would land outside the parens — use P (paste before) to drop " +
					"the value back inside where the ) is waiting.",
				NewKeys:     []string{"yi("},
				Start:       []string{"scale(2.0)", "scale(1.0)"},
				CursorStart: Pos{0, 6},
				Kind:        KindEdit,
				Target:      []string{"scale(2.0)", "scale(2.0)"},
				Par:         10,
			},
			{
				Title: "Copy The Name",
				Instructions: "One more delimiter, same three-step order: delete the wrong " +
					"quoted name first with di\", yank the right one from the line above " +
					"with yi\", then paste it back in with P.",
				Tip: "Tip: di\" leaves the cursor on the closing quote too, just like di( " +
					"did on the closing paren — pasting before the cursor with P is what " +
					"drops the text back inside the pair.",
				NewKeys:     []string{"yi\""},
				Start:       []string{`name = "Alice"`, `name = "Bob"`},
				CursorStart: Pos{0, 8},
				Kind:        KindEdit,
				Target:      []string{`name = "Alice"`, `name = "Alice"`},
				Par:         10,
			},
			{
				Title: "Capstone",
				Instructions: "One template line on top, three broken copies below it — " +
					"one wrong word, one wrong number, one wrong path. Fix all three the " +
					"same way: delete the wrong piece first, jump back up to the template " +
					"to yank the correct one, then return and paste it in.",
				Tip: "Tip: re-locate with 0 then the same motion you used the first time " +
					"(w w, or f( / f\") before every delete and every paste — don't rely on " +
					"the cursor's column carrying over between lines, since yanking and " +
					"deleting don't update it the way real motions do.",
				Start: []string{
					`role = admin, count(5), path = "main.go"`,
					`role = guest, count(5), path = "main.go"`,
					`role = admin, count(0), path = "main.go"`,
					`role = admin, count(5), path = "temp.go"`,
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					`role = admin, count(5), path = "main.go"`,
					`role = admin, count(5), path = "main.go"`,
					`role = admin, count(5), path = "main.go"`,
					`role = admin, count(5), path = "main.go"`,
				},
				Par: 63,
			},
		},
	},
	{
		Number: 20,
		Week:   "Week 3: Counts & Text Objects",
		Title:  "Y — Yank Whole Line, The Fast Way",
		Summary: "Y is a shortcut for yy — one keystroke instead of two to yank the whole " +
			"current line — and it takes a count exactly the same way yy does, so 3Y and " +
			"3yy are identical. The real payoff shows up when you combine it with the text " +
			"objects you already know: yank a line as a template with Y, paste a copy of it " +
			"with p, then tweak just the one word, number, or quoted value that's different " +
			"with ciw, ci(, or ci\" — a duplicate-and-customize pattern you'll use constantly.",
		Challenges: []Challenge{
			{
				Title: "Yank The Whole Line, Fast",
				Instructions: "Y does exactly what yy does — yanks the whole current line — " +
					"in one keystroke instead of two. Yank this line with Y, then paste a " +
					"copy of it below with p.",
				Tip: "Tip: Y and yy are the same command in this engine; use whichever " +
					"your fingers reach for.",
				NewKeys:     []string{"Y"},
				Start:       []string{"keep me twice"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"keep me twice", "keep me twice"},
				Par:         2,
			},
			{
				Title: "Count Before Y",
				Instructions: "Y takes a count exactly like yy does — 2Y yanks two lines " +
					"starting from the cursor in a single command. Copy the two template " +
					"lines with 2Y, jump to the bottom with G, then paste the block after " +
					"the last line with p.",
				Tip: "Tip: 2Y and 2yy do the exact same thing — the count trick from Day " +
					"15 works on every one of these shortcuts, not just yy.",
				Start:       []string{"template A", "template B", "keep me", "tail"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"template A", "template B", "keep me", "tail", "template A", "template B",
				},
				Par: 4,
			},
			{
				Title: "Duplicate And Customize",
				Instructions: "Copy a line as a starting point, then tweak just one word in " +
					"the copy — a very common real edit. Yank this line with Y, paste it " +
					"below with p, then use ciw on the pasted copy to change 'draft' to " +
					"'final'.",
				Tip: "Tip: after p, your cursor lands right at the start of the pasted " +
					"line — no need to jump down to it first.",
				Start:       []string{"status: draft"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"status: draft", "status: final"},
				Par:         14,
			},
			{
				Title: "Duplicate And Customize, Again",
				Instructions: "Same pattern, this time on a number inside parens: Y and p " +
					"to duplicate the line, then ci( on the copy to change 1 to 5.",
				Tip: "Tip: this is the same Y, p, then-fix-the-copy rhythm as the last " +
					"challenge — only the delimiter you edit with changes.",
				Start:       []string{"retries(1)"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"retries(1)", "retries(5)"},
				Par:         10,
			},
			{
				Title: "Capstone",
				Instructions: "A three-line settings block. Yank all three at once with 3Y, " +
					"jump to the bottom with G, and paste the whole block with p. Then fix " +
					"the three copies one at a time: ci\" to change the quoted name to " +
					"prod, ci( to change the count to 9, and ciw to change the role to " +
					"admin.",
				Tip: "Tip: this is every trick from Days 19 and 20 in one document — yank " +
					"a whole block with a count, paste it, then reach into the copy with " +
					"text objects to fix exactly the pieces that need to change.",
				Start: []string{
					`name = "template"`,
					`count(1)`,
					`role = user`,
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					`name = "template"`,
					`count(1)`,
					`role = user`,
					`name = "prod"`,
					`count(9)`,
					`role = admin`,
				},
				Par: 37,
			},
		},
	},
	{
		Number: 21,
		Week:   "Week 3: Full Revision",
		Title:  "Full Revision — The Ultimate Boss (Everything From Days 1-20)",
		Summary: "No new keys today — this is the hardest day yet, and it's built to prove " +
			"you actually own everything from the last three weeks, not just recognize it. " +
			"Five challenges, each pulling from a different stretch of the curriculum: basic " +
			"movement and line ops, word operators and text objects, replace/substitute/undo " +
			"and the line-position motions, counts with every bracket-and-quote text object " +
			"plus yanking them, and a genuinely long final-boss document that mixes all of it " +
			"together in one continuous pass. Take your time — this is meant to feel like a " +
			"real workout, not a quick clear.",
		Challenges: []Challenge{
			{
				Title: "Movement, Insert & Line Ops",
				Instructions: "Five lines, only Week 1 tools. Fix the stray digit in the " +
					"first line with l, x, and i. Skip the second line completely — not " +
					"every line needs touching. Delete the third line outright with dd. " +
					"Duplicate the fourth line below itself with yy and p. Then jump to " +
					"the very last line with G and tack a suffix on with A.",
				Tip: "Tip: nothing here is new — hjkl, i/a/o/O, gg/G/x, and dd/yy/p/P are " +
					"exactly what they were in Week 1. The only new part is chaining five " +
					"of them without stopping to think about each one in isolation.",
				Start: []string{
					"wor1d needs a fix",
					"jump target line stays put",
					"extra line to remove completely",
					"footer text to duplicate",
					"last line needs a suffix",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"world needs a fix",
					"jump target line stays put",
					"footer text to duplicate",
					"footer text to duplicate",
					"last line needs a suffix - done",
				},
				Par: 24,
			},
			{
				Title: "Word Operators & Text Objects",
				Instructions: "Seven lines, one job each. dw removes a word plus its " +
					"trailing space. cw only touches the word itself — the classic quirk " +
					"from Day 6. ciw is the explicit version of that same quirk. diw " +
					"deletes just the word and leaves the surrounding whitespace alone. " +
					"daw cleans up the whitespace too. D cuts to the end of the line from " +
					"wherever the cursor sits, and C does the same but drops you into " +
					"Insert mode to type a replacement.",
				Tip: "Tip: diw vs daw is the detail people forget fastest — diw leaves a " +
					"gap behind, daw doesn't. If a line ends up with a double space after " +
					"an edit, that's diw working exactly as designed.",
				Start: []string{
					"quick brown extra fox jumps",
					"typo old word stays put",
					"clean inner word please fix",
					"remove inner gap word please",
					"trim around this word right here",
					"cut this tail right after here",
					"replace this tail right after here",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"quick brown fox jumps",
					"typo new word stays put",
					"clean outer word please fix",
					"remove inner  word please",
					"trim around word right here",
					"cut this tail",
					"replace this tail section",
				},
				Par: 62,
			},
			{
				Title: "Replace, Substitute & Line Motions",
				Instructions: "Nine lines pulling from Days 9-13: r to swap one character, " +
					"~ to flip case a run at a time, s to substitute a character (with a " +
					"typed replacement), S to rewrite a whole line, a counted X to delete " +
					"backward through a garbage prefix, u to undo a wrong delete before " +
					"redoing it correctly, and 0/^/$ to prepend a note, cut trailing junk, " +
					"and replace an indented prefix.",
				Tip: "Tip: undo reverts one whole action, not one keystroke — so u after a " +
					"multi-character insert or a whole dw undoes all of it in a single " +
					"press, exactly like it did back on Day 12.",
				Start: []string{
					"cat sat on the mit",
					"hello WORLD",
					"fix 5 items today",
					"old placeholder line",
					"###keep this",
					"alpha beta gamma delta",
					"    fix the indentation on this line",
					"keep this good part TRASH from here onward",
					"    XXXX apply the fix here too",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"cat sat on the mat",
					"HELLO world",
					"fix five items today",
					"brand new line content",
					"keep this",
					"beta gamma delta",
					"NOTE:     fix the indentation on this line",
					"keep this good part ",
					"    DONE: apply the fix here too",
				},
				Par: 101,
			},
			{
				Title: "Counts, Every Text Object & Y",
				Instructions: "Everything from Days 15-20 in one pass: 3dd to delete a " +
					"block with a count, then all five bracket-and-quote text objects — " +
					"i(, i[, i{, and i' — each fixing a value inside its pair, then three " +
					"yank-and-fix pairs using yiw, yi(, and yi\" to copy a correct value " +
					"from one line into a broken one below it, and finally Y and p to " +
					"duplicate a template line the fast way.",
				Tip: "Tip: y never touches the buffer — on the yank/paste pairs, delete " +
					"the wrong value first, yank the right one second, and paste last, or " +
					"the delete will overwrite whatever you just yanked.",
				Start: []string{
					"skip one",
					"skip two",
					"skip three",
					"keep this one",
					"call(oldArg, extra)",
					"arr[0] stays",
					"{debug=false}",
					"name is 'Alice' today",
					"status = active",
					"status = pending",
					"scale(2.0)",
					"scale(1.0)",
					`name = "Alice"`,
					`name = "Bob"`,
					"keep me twice",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"keep this one",
					"call(newArg)",
					"arr[42] stays",
					"{debug=true}",
					"name is 'Bob' today",
					"status = active",
					"status = active",
					"scale(2.0)",
					"scale(2.0)",
					`name = "Alice"`,
					`name = "Alice"`,
					"keep me twice",
					"keep me twice",
				},
				Par: 104,
			},
			{
				Title: "Final Boss — Everything From Days 1-20",
				Instructions: "One long document, fourteen lines, every tool from the last " +
					"three weeks back to back: a garbage prefix to replace with c^, a " +
					"duplicated line to remove with dd, a typo to fix with ciw, an entire " +
					"junk line to delete, a number inside parens to redact with ci(, a " +
					"shouted line to fix with repeated ~, a trailing phrase to trim with " +
					"D, three lines to reorder with dd/P, a single-quoted value to change " +
					"with ci', a bracketed index to change with ci[, a template line to " +
					"duplicate and customize with Y/p/ciw, and a brand new closing line to " +
					"add at the very end with G and o. Nothing here is new — it's every " +
					"single tool from Days 1-20, on a document big enough to actually take " +
					"real time to clear properly.",
				Tip: "Tip: work top to bottom and don't rush. If you get stuck on any one " +
					"line, it's testing exactly one thing from one specific earlier day — " +
					"think back to which day taught that shape of problem and the tool " +
					"will come back to you.",
				Start: []string{
					"    XXXX quarterly report begins here",
					"the numbers below need review",
					"the numbers below need review",
					"we fond a typo in this sentence",
					"delete this entire line of scratch notes",
					"budget(999) needs redacting",
					"STOP SHOUTING in this summary",
					"trim the tail needless extra words here",
					"cherry",
					"banana",
					"apple",
					"config = 'debug'",
					"list[0] = old",
					"status: draft",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"    DONE: quarterly report begins here",
					"the numbers below need review",
					"we found a typo in this sentence",
					"budget(XXX) needs redacting",
					"stop shouting in this summary",
					"trim the tail",
					"apple",
					"cherry",
					"banana",
					"config = 'release'",
					"list[1] = old",
					"status: draft",
					"status: final",
					"report complete",
				},
				Par: 128,
			},
		},
	},
	{
		Number: 22,
		Week:   "Week 4: Find & Search",
		Title:  "f, F, t, T — Jump Straight To Any Character",
		Summary: "A whole new family of motions: instead of stepping one cell at a time, " +
			"jump straight to a character anywhere on the current line. f{char} jumps " +
			"forward and lands ON the character. F{char} jumps backward and lands ON it " +
			"too. t{char} jumps forward and stops just BEFORE it. T{char} jumps backward " +
			"and stops just AFTER it. Then ; repeats the last find in the same direction, " +
			"and , repeats it in the opposite direction — so you almost never retype the " +
			"character. These combine with operators exactly like any other motion: df, " +
			"dt, cf, ct all work.",
		Challenges: []Challenge{
			{
				Title: "f — Jump Forward, Land On The Character",
				Instructions: "f{char} scans forward on this line and drops the cursor " +
					"exactly on the next occurrence of that character. Press f then o to " +
					"land on the o in \"box\" — then press f then o again to jump past it " +
					"to the next o, in \"dozen\". That's the goal.",
				Tip: "Tip: f doesn't wrap to the next line and doesn't care about words — " +
					"it only cares about the exact character you give it, anywhere ahead " +
					"of the cursor on the current line.",
				Start: []string{
					"pack my box with five dozen liquor jugs",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 23},
				Par:         4,
			},
			{
				Title: "F and T — Jump Backward, On It Or Just After It",
				Instructions: "F{char} is f in reverse: scan backward, land ON the " +
					"character. T{char} also scans backward, but stops just AFTER it — " +
					"the mirror image of t. Starting from the end of this line, press F " +
					"then j to land on the j in \"jumps\", then T then t to jump further " +
					"back and stop right after the t in \"the\".",
				Tip: "Tip: T stopping just after the character (instead of on it) is " +
					"exactly like t, just facing the other direction — same rule, " +
					"opposite scan direction.",
				Start: []string{
					"the quick brown fox jumps over the lazy dog",
				},
				CursorStart: Pos{0, 42},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 1},
				Par:         4,
			},
			{
				Title: "dt, df, ct, cf — Find Meets An Operator",
				Instructions: "Any operator can take a find-motion instead of a plain " +
					"motion. On line 1, ct: changes everything up to (not including) the " +
					"colon — the colon survives, exactly like ciw vs diw all over again. " +
					"On line 2, df: deletes everything up to AND including the colon in " +
					"one shot — no separate x needed to clean up the delimiter.",
				Tip: "Tip: t is exclusive (stops short, the boundary character survives), " +
					"f is inclusive (eats the boundary character too) — same distinction " +
					"as di( vs the character itself, just for find-motions now.",
				Start: []string{
					"URGENT: fix this bug",
					"DELETE-ME:keep this text",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"note: fix this bug",
					"keep this text",
				},
				Par: 13,
			},
			{
				Title: "; and , — Repeat The Last Find",
				Instructions: "; repeats your last f/F/t/T in the same direction, and , " +
					"repeats it in the opposite direction — no retyping the character. " +
					"On line 1, use f, then ; once to hop two commas over and land right " +
					"before \"blue\", then fix it. On line 2, use f, then ; twice to " +
					"overshoot to the comma before \"date\", then , once to step back to " +
					"the comma before \"cherry\", then fix that instead.",
				Tip: "Tip: , is perfect for exactly this — you overshoot by one while " +
					"chaining ;, so instead of starting the whole search over you just " +
					"tap , once to back up a single step.",
				Start: []string{
					"red,green,blue,yellow,purple",
					"apple,banana,cherry,date,fig",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"red,green,black,yellow,purple",
					"apple,banana,grape,date,fig",
				},
				Par: 30,
			},
			{
				Title: "Find Everything — A Log File Full Of Fixes",
				Instructions: "Five lines, five different find-motion tricks. Line 1: " +
					"jump to the swapped i with f, then fix the recieve/receive typo " +
					"with the classic xp swap. Line 2: jump to the P in PENDING with f, " +
					"then ct: to rename it to DONE. Line 3: df with a space deletes the " +
					"whole [DEBUG] tag in one inclusive find. Line 4: from the end of the " +
					"line, F back to the W in Warn and fix it with ciw. Line 5: f: then " +
					"; twice to reach the third field, then ciw to fix its value.",
				Tip: "Tip: nothing here is a new key — it's f, F, t, T, and ; doing the " +
					"legwork so that x, p, ciw, and c can land exactly where they need to " +
					"without a single wasted hjkl step.",
				Start: []string{
					"recieve the shipment confirmation first",
					"STATUS-PENDING: awaiting review",
					"[DEBUG] connection established",
					"log entry: level=Warn source=Legacy-Module",
					"CPU:12,MEM:45,DISK:78,NET:33",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"receive the shipment confirmation first",
					"STATUS-DONE: awaiting review",
					"connection established",
					"log entry: level=Error source=Legacy-Module",
					"CPU:12,MEM:45,DISK:90,NET:33",
				},
				Par: 47,
			},
		},
	},
	{
		Number: 23,
		Week:   "Week 4: Find & Search",
		Title:  "/ and ? — Search The Whole Buffer",
		Summary: "f and t only look at the current line. / and ? look at the whole " +
			"buffer. Type / followed by a word and press enter to jump forward to the " +
			"next place that word appears, anywhere in the document. ? does the same " +
			"thing backward. Once you've searched, n repeats the same search in the same " +
			"direction, and N repeats it in the opposite direction — just like ; and , " +
			"did for f and t.",
		Challenges: []Challenge{
			{
				Title: "/ — Search Forward",
				Instructions: "Press / to open a search prompt at the bottom, type ERROR, " +
					"and press enter. The cursor jumps straight to it, no matter how many " +
					"lines away it is.",
				Tip: "Tip: unlike f/t, / isn't limited to the current line — it scans " +
					"every line in the buffer for the next match.",
				Start: []string{
					"connecting to server...",
					"loading configuration",
					"validating credentials",
					"ERROR: connection timeout",
					"retrying in 5 seconds",
					"connection established",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{3, 0},
				Par:         7,
			},
			{
				Title: "? — Search Backward",
				Instructions: "Same idea, opposite direction. Starting from the bottom of " +
					"this list, press ? followed by configure and enter to jump backward " +
					"to the line that mentions it.",
				Tip: "Tip: ? is to / what F is to f — same search, just scanning upward " +
					"through the buffer instead of downward.",
				Start: []string{
					"step one: prepare",
					"step two: configure",
					"step three: deploy",
					"step four: verify",
					"step five: cleanup",
				},
				CursorStart: Pos{4, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{1, 10},
				Par:         11,
			},
			{
				Title: "n — Repeat The Last Search",
				Instructions: "Search for pending, then press n twice to skip past the " +
					"first two matches and land on the third one — the one on its own " +
					"line, not the one buried in the summary. Fix it with ciw.",
				Tip: "Tip: n always repeats in the same direction the search was made, " +
					"whether that search was / or ? — no need to remember which one you " +
					"used.",
				Start: []string{
					"task pending",
					"task pending",
					"task done",
					"task pending",
					"summary: 3 tasks pending, 1 done",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"task pending",
					"task pending",
					"task done",
					"task complete",
					"summary: 3 tasks pending, 1 done",
				},
				Par: 23,
			},
			{
				Title: "N — Repeat In The Opposite Direction",
				Instructions: "Starting from the bottom, search backward for ok, then " +
					"press n twice more to keep walking backward. That overshoots past " +
					"server-c. Press N once to reverse direction and step back onto it, " +
					"then fix it with ciw.",
				Tip: "Tip: N is the escape hatch for exactly this — walk too far with n " +
					"and a single N walks back one step without restarting the search.",
				Start: []string{
					"server-a: status ok",
					"server-b: status ok",
					"server-c: status ok",
					"server-d: status ok",
					"server-e: status ok",
				},
				CursorStart: Pos{4, 0},
				Kind:        KindEdit,
				Target: []string{
					"server-a: status ok",
					"server-b: status ok",
					"server-c: status down",
					"server-d: status ok",
					"server-e: status ok",
				},
				Par: 15,
			},
			{
				Title: "A Fleet Of Alerts",
				Instructions: "Six servers, three still saying alert. Search forward for " +
					"alert and fix node-2 to warning. Press n to reach node-4 and fix it " +
					"to critical. Then search backward for alert to catch node-1, fixing " +
					"it to resolved. Finally press N to reverse back to forward and catch " +
					"node-6, fixing it to monitoring.",
				Tip: "Tip: once you fix a line, its old text stops matching the search — " +
					"so n and N never land you back on a line you already handled.",
				Start: []string{
					"node-1: alert - disk space low",
					"node-2: alert - disk space low",
					"node-3: status ok",
					"node-4: alert - disk space low",
					"node-5: status ok",
					"node-6: alert - disk space low",
				},
				CursorStart: Pos{0, 13},
				Kind:        KindEdit,
				Target: []string{
					"node-1: resolved - disk space low",
					"node-2: warning - disk space low",
					"node-3: status ok",
					"node-4: critical - disk space low",
					"node-5: status ok",
					"node-6: monitoring - disk space low",
				},
				Par: 65,
			},
		},
	},
	{
		Number: 24,
		Week:   "Week 4: Find & Search",
		Title:  "% — Jump To The Matching Bracket",
		Summary: "Code is full of ( ), [ ], and { } pairs, and they're often several " +
			"words or even several lines apart. Put the cursor on (or just before) any " +
			"bracket and press % to jump straight to its match — the other end of the " +
			"same pair. It works across multiple lines, and it correctly skips over any " +
			"nested pairs of the same bracket type along the way.",
		Challenges: []Challenge{
			{
				Title: "% — Jump To The Closer",
				Instructions: "Your cursor sits on the ( of calculate(...). Press % and " +
					"watch the cursor jump straight to the matching ) at the end of the " +
					"line.",
				Tip: "Tip: % is the one motion that leaps across however much text sits " +
					"between a bracket and its partner, no counting required.",
				Start:       []string{"calculate(alpha, beta, gamma)"},
				CursorStart: Pos{0, 9},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 28},
				Par:         1,
			},
			{
				Title: "% Crosses Lines Too",
				Instructions: "This function body spans four lines. Your cursor is on the " +
					"{ that opens it. Press % to jump all the way down to the closing } " +
					"on its own line.",
				Tip: "Tip: unlike f/t, % isn't limited to the current line — same as / " +
					"and ?, it can land anywhere in the buffer.",
				Start: []string{
					"function setup() {",
					"    initialize()",
					"    configure()",
					"    validate()",
					"}",
				},
				CursorStart: Pos{0, 17},
				Kind:        KindGoal,
				GoalPos:     Pos{4, 0},
				Par:         1,
			},
			{
				Title: "d% — Delete A Bracket Pair And Everything Inside",
				Instructions: "Your cursor is on the ( of run(...). Press d% to delete " +
					"from that ( through the matching ) — brackets and all — leaving " +
					"just run behind.",
				Tip: "Tip: % combines with operators cleanly, same as f/t — d% deletes " +
					"the whole bracketed chunk in one move, no matter how long it is.",
				Start:       []string{"run(oldTask, extra, args)", "keep this line unchanged"},
				CursorStart: Pos{0, 3},
				Kind:        KindEdit,
				Target:      []string{"run", "keep this line unchanged"},
				Par:         2,
			},
			{
				Title: "% Skips Nested Brackets",
				Instructions: "This array has a smaller array nested inside it. Your " +
					"cursor is on the outer [. Press c% and type REDACTED to replace " +
					"the whole outer list — % correctly jumps past the inner [20, 30] " +
					"to the real outer ], so nothing gets left behind.",
				Tip: "Tip: % counts nesting depth as it scans, so it never mistakes an " +
					"inner closing bracket for the outer one you actually meant.",
				Start:       []string{"values = [10, [20, 30], 40]", "next line unchanged"},
				CursorStart: Pos{0, 9},
				Kind:        KindEdit,
				Target:      []string{"values = REDACTED", "next line unchanged"},
				Par:         11,
			},
			{
				Title: "Tidy The Settings Block",
				Instructions: "Fix three values using % to reach each one, then confirm " +
					"the whole block's extent. Line 2: jump to the ( in timeout and use " +
					"c% to replace (30 + 5) with (45). Line 3: jump to the [ in retries " +
					"and use c% to replace [1, 2, 3] with [5, 10]. Line 4: jump to the " +
					"opening quote in name and use ci\" to change prod to staging. " +
					"Finally, go back to the very first line, press $ to reach its { , " +
					"then % to confirm it matches the closing } at the bottom, and press " +
					"o to add a new line after it reading # end of settings.",
				Tip: "Tip: % is as much a sanity check as an edit tool — jumping from an " +
					"opening brace to its close confirms exactly how much a block spans.",
				Start: []string{
					"settings = {",
					"    timeout: (30 + 5),",
					"    retries: [1, 2, 3],",
					"    name: \"prod\"",
					"}",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"settings = {",
					"    timeout: (45),",
					"    retries: [5, 10],",
					"    name: \"staging\"",
					"}",
					"# end of settings",
				},
				Par: 63,
			},
		},
	},
	{
		Number: 25,
		Week:   "Week 4: Find & Search",
		Title:  "* and # — Search The Word Under The Cursor",
		Summary: "With / you had to type the word you were looking for. * skips the " +
			"typing entirely: put the cursor anywhere on a word and press * to jump to " +
			"its next occurrence, or # to jump to its previous one. Both match the " +
			"whole word only — searching for count will never land inside counter — " +
			"and n/N repeat the jump exactly like they do for / and ?.",
		Challenges: []Challenge{
			{
				Title: "* — Jump To The Next Occurrence",
				Instructions: "Your cursor sits in the middle of total on the first " +
					"line. Press * and jump straight to the next place total appears — " +
					"no typing, no search prompt.",
				Tip: "Tip: the cursor can be anywhere inside the word — * figures out " +
					"the word's boundaries for you before it searches.",
				Start: []string{
					"total = price + tax",
					"shipping = 5",
					"final = total - discount",
				},
				CursorStart: Pos{0, 2},
				Kind:        KindGoal,
				GoalPos:     Pos{2, 8},
				Par:         1,
			},
			{
				Title: "# — Search Backward, Whole Words Only",
				Instructions: "Your cursor is on count in the last line. Press # to " +
					"jump backward to the previous count — notice it skips right over " +
					"counter on the way. Then press n to keep going backward to the " +
					"definition at the very top.",
				Tip: "Tip: * and # match only the whole word. A plain /count search " +
					"would have stopped inside counter; # never will.",
				Start: []string{
					"count = 0",
					"counter = count + 1",
					"update(counter)",
					"print(count)",
				},
				CursorStart: Pos{3, 6},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 0},
				Par:         2,
			},
			{
				Title: "* Then Fix It",
				Instructions: "Both function calls should use the local limit variable, " +
					"not the global timeout. Your cursor is on timeout in the " +
					"definition. Press * to jump to the one inside retry and fix it " +
					"with ciw limit, then n to reach the one inside connect and fix it " +
					"the same way.",
				Tip: "Tip: same rule as / — search is a pure jump in this editor, so " +
					"land first with * or n, then edit with a separate command like ciw.",
				Start: []string{
					"timeout = 30",
					"retry(timeout, backoff)",
					"connect(host, timeout)",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"timeout = 30",
					"retry(limit, backoff)",
					"connect(host, limit)",
				},
				Par: 20,
			},
			{
				Title: "n Wraps Around — Rename Every Occurrence",
				Instructions: "Rename tmp to data everywhere. From the definition, " +
					"press * to reach the second occurrence and fix it with ciw data, " +
					"press n and fix the third, then press n once more — the search " +
					"wraps past the end of the file, straight back to the first tmp — " +
					"and fix that one too.",
				Tip: "Tip: wrap-around is what makes *-then-n a full rename loop — " +
					"start anywhere, keep pressing n, and you'll visit every occurrence " +
					"exactly once.",
				Start: []string{
					"tmp = load()",
					"result = tmp + 2",
					"save(tmp)",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"data = load()",
					"result = data + 2",
					"save(data)",
				},
				Par: 27,
			},
			{
				Title: "Clean Up The Connection Script",
				Instructions: "Three jobs. First: your cursor is on retries at the " +
					"bottom — press # to jump back to its definition, then use $ and " +
					"r to change the 3 to a 5. Second: press j then 0 to land on tmp, " +
					"and rename every tmp to conn using * and n (the last n wraps back " +
					"to the definition) with ciw conn each time. Third: press G, then " +
					"o, and add a final line reading done.",
				Tip: "Tip: # to find where something was defined, * to sweep every " +
					"use of it — together they cover both directions of the rename " +
					"workflow.",
				Start: []string{
					"retries = 3",
					"tmp = connect()",
					"send(tmp)",
					"close(tmp)",
					"check(retries)",
				},
				CursorStart: Pos{4, 6},
				Kind:        KindEdit,
				Target: []string{
					"retries = 5",
					"conn = connect()",
					"send(conn)",
					"close(conn)",
					"check(retries)",
					"done",
				},
				Par: 40,
			},
		},
	},
	{
		Number: 26,
		Week:   "Week 4: Find & Search",
		Title:  "d/, c/, y/ — Operators Meet Search",
		Summary: "/ and ? aren't just for moving the cursor — any operator can take a " +
			"search as its motion, exactly like it already takes f, t, or %. d/{pattern} " +
			"then enter deletes from the cursor up to (not including) the next match, " +
			"c/{pattern} does the same and drops you into insert mode, and y/{pattern} " +
			"copies the span into the register without touching the buffer. ? works the " +
			"same way in reverse, deleting or yanking backward to the previous match " +
			"instead. Just like d%, the range can cross as many lines as it needs to.",
		Challenges: []Challenge{
			{
				Title: "d/ — Delete Up To The Next Match",
				Instructions: "d/{pattern} then enter deletes everything between the " +
					"cursor and the next match, stopping just short of it — the match " +
					"itself survives. Press d then / then type fix and enter to strip " +
					"the noisy prefix off this line.",
				Tip: "Tip: this is exclusive, same rule as dt or dw — whatever you " +
					"search for is the boundary, and boundaries you find (rather than " +
					"typing a character for) still follow it.",
				Start:       []string{"URGENT: fix this bug now"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"fix this bug now"},
				Par:         6,
			},
			{
				Title: "d? — Delete Backward To A Match",
				Instructions: "? searches backward, so d?{pattern} deletes backward " +
					"instead — from the match through everything up to (not including) " +
					"where your cursor started. Your cursor sits on \"this\". Press d " +
					"then ? then type TODO and enter to remove the marker and the " +
					"clutter after it in one shot.",
				Tip: "Tip: d? is the mirror of d/ — the match itself gets swallowed " +
					"this time because it's now the far end of the range instead of " +
					"the boundary you stop before.",
				Start:       []string{"remove TODO cleanup this stray text here"},
				CursorStart: Pos{0, 20},
				Kind:        KindEdit,
				Target:      []string{"remove this stray text here"},
				Par:         7,
			},
			{
				Title: "c/ — Change Up To A Match",
				Instructions: "c/ behaves exactly like c does with any other motion: " +
					"delete the exclusive range, then drop into insert mode to type " +
					"the replacement. Press c then / then type PENDING and enter to " +
					"reach the status flag, then type DONE followed by a space so it " +
					"still reads cleanly against what comes after.",
				Tip: "Tip: same exclusive-boundary idea as ct: from day 22 — the text " +
					"you searched for stays put, only what's in front of it changes.",
				Start:       []string{"status: PENDING please wait"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"DONE PENDING please wait"},
				Par:         16,
			},
			{
				Title: "y/ — Yank Up To A Match, Then Paste It Elsewhere",
				Instructions: "y/ never touches the buffer — it just fills the " +
					"register with everything up to the match, same as y always has. " +
					"Your cursor sits right after \"endpoint: \" on line 1. Press y " +
					"then / then type method and enter to copy the path, then press j " +
					"and P to drop it into the gap on line 2, right before its own " +
					"\"method\".",
				Tip: "Tip: the search that y/ uses to find its boundary doesn't move " +
					"your cursor any further than the boundary itself — same as every " +
					"other y motion, it leaves the buffer alone until you paste.",
				Start: []string{
					"endpoint: /api/v1/users method=GET",
					"endpoint: method=POST",
				},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target: []string{
					"endpoint: /api/v1/users method=GET",
					"endpoint: /api/v1/users method=POST",
				},
				Par: 11,
			},
			{
				Title: "Tidy The Startup Log",
				Instructions: "Five lines, four search-powered fixes. Line 1: " +
					"d/server clears the DEBUG prefix. Line 2: f= then l lands you " +
					"right after the equals sign, then c/code changes PENDING to DONE " +
					"(type DONE with a trailing space). Line 3 to line 4: w w reaches " +
					"retries=5, y/backoff copies it, then j and P paste it into line " +
					"4's gap. Line 5: w w w reaches flag, then d?TODO deletes the " +
					"marker and the clutter behind it, leaving flag here in place.",
				Tip: "Tip: nothing here is a new key — it's / and ? doing the same job " +
					"f and t did back on day 22, just aimed at the whole buffer instead " +
					"of a single line.",
				Start: []string{
					"DEBUG: server booting up",
					"STATUS=PENDING code=200",
					"template: retries=5 backoff=200",
					"override: backoff=200",
					"cleanup TODO leftover flag here",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"server booting up",
					"STATUS=DONE code=200",
					"template: retries=5 backoff=200",
					"override: retries=5 backoff=200",
					"cleanup flag here",
				},
				Par: 55,
			},
		},
	},
	{
		Number: 27,
		Week:   "Week 4: Find & Search",
		Title:  "Search Patterns — \\d, ^, and Character Classes",
		Summary: "/ and ? don't just match literal text — they run real search " +
			"patterns. \\d matches any digit, so /\\d jumps to the next one without " +
			"you typing which digit it is. ^ anchors a pattern to the start of a " +
			"line, so /^word only matches word when it opens the line, not when it " +
			"shows up in the middle of one. [set] matches any single character from " +
			"the set inside the brackets — [,;] matches a comma or a semicolon, " +
			"whichever comes first. These patterns work anywhere / or ? already " +
			"worked, including n, N, and the d/ c/ y/ operators from day 26.",
		Challenges: []Challenge{
			{
				Title: "\\d — Match Any Digit",
				Instructions: "\\d stands in for any single digit, 0 through 9 — " +
					"you don't have to know which one shows up first. Press / then " +
					"\\d and enter to jump straight to the first digit in this line.",
				Tip: "Tip: \\d is a class, not a literal backslash-d — same idea as " +
					"typing a real character for f, except this one matches a whole " +
					"category of characters instead of just one.",
				Start:       []string{"invoice INV-2024 total due"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 12},
				Par:         4,
			},
			{
				Title: "^ — Anchor To The Start Of A Line",
				Instructions: "^ pins a pattern to the very start of a line, so it " +
					"only matches there — not anywhere else the text happens to " +
					"appear. Press / then ^WARN and enter to skip past the WARN " +
					"buried mid-line and land on the one line that actually starts " +
					"with it.",
				Tip: "Tip: without the ^, /WARN would have stopped on the very " +
					"first line instead — WARN appears there too, just not at the " +
					"front.",
				Start: []string{
					"status: WARN check disk",
					"WARN: disk usage high",
					"note: ignore WARN here",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{1, 0},
				Par:         7,
			},
			{
				Title: "[set] — Match Any One Of Several Characters",
				Instructions: "[,;] matches a single character that's either a " +
					"comma or a semicolon — whichever one the search hits first. " +
					"Press / then [,;] and enter to land on the first separator in " +
					"this line.",
				Tip: "Tip: put any characters you want to match inside the " +
					"brackets — [,;] means \"comma or semicolon\", not the literal " +
					"three-character string.",
				Start:       []string{"field_one; field_two, field_three"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 9},
				Par:         6,
			},
		},
	},
}

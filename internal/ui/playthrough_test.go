package ui

import "testing"

func TestCurriculumSolvable(t *testing.T) {
	cases := []struct {
		day, challenge int
		keys           []string
	}{
		{1, 0, []string{"j", "j", "j", "l", "l", "l", "l"}},
		{1, 1, []string{"j", "j", "j", "j", "j", "j", "l", "l", "l", "l", "l", "l", "l"}},
		{1, 2, []string{"j", "j", "j", "j", "j", "j", "j", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l"}},
		{1, 3, []string{"k", "k", "k", "k", "k", "h", "h", "h", "h", "h"}},
		{1, 4, []string{"k", "k", "k", "k", "k", "k", "k", "k", "k", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l"}},

		{2, 0, []string{"w", "w", "w", "w", "w", "w", "w", "w"}},
		{2, 1, []string{"b", "b", "b", "b", "b", "b"}},
		{2, 2, []string{"e", "e", "e", "e", "e", "e", "e", "e", "e"}},
		{2, 3, []string{"w", "w", "w", "w", "w", "w", "w", "w"}},
		{2, 4, []string{"b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}},

		{3, 0, append([]string{"i"}, append(typeStr("welcome to vim insert mode"), "esc")...)},
		{3, 1, append([]string{"A"}, append(typeStr(" world of vim editing"), "esc")...)},
		{3, 2, append([]string{"o"}, append(typeStr("second line goes right here"), "esc")...)},
		{3, 3, append([]string{"O"}, append(typeStr("first line"), "esc")...)},
		{3, 4, concatKeys(
			append([]string{"O"}, append(typeStr("appetizer"), "esc")...),
			[]string{"j"},
			append([]string{"A"}, append(typeStr(" dish"), "esc")...),
			append([]string{"o"}, append(typeStr("dessert"), "esc")...),
		)},

		{4, 0, []string{"g", "g"}},
		{4, 1, []string{"G"}},
		{4, 2, []string{"x", "j", "x", "j", "x"}},
		{4, 3, []string{"7", "g", "g"}},
		{4, 4, []string{"k", "k", "x", "j", "j", "j", "x", "G", "x"}},

		{5, 0, []string{"j", "d", "d", "j", "j", "d", "d"}},
		{5, 1, []string{"k", "y", "y", "p", "p"}},
		{5, 2, []string{"y", "y", "P", "P"}},
		{5, 3, []string{"d", "d", "G", "p"}},
		{5, 4, []string{"j", "d", "d", "p", "G", "d", "d", "k", "p"}},

		{6, 0, []string{"d", "w", "d", "w"}},
		{6, 1, concatKeys(
			append([]string{"c", "w"}, append(typeStr("quick"), "esc")...),
			[]string{"w"},
			append([]string{"c", "w"}, append(typeStr("brown"), "esc")...),
		)},
		{6, 2, concatKeys(
			append([]string{"c", "i", "w"}, append(typeStr("word"), "esc")...),
			[]string{"w", "w"},
			append([]string{"c", "i", "w"}, append(typeStr("time"), "esc")...),
		)},
		{6, 3, []string{"d", "i", "w"}},
		{6, 4, concatKeys(
			[]string{"d", "w"},
			[]string{"j", "w"},
			append([]string{"c", "w"}, append(typeStr("quick"), "esc")...),
			[]string{"j", "b"},
			append([]string{"c", "i", "w"}, append(typeStr("word"), "esc")...),
		)},

		{7, 0, []string{"d", "a", "w", "w", "d", "a", "w"}},
		{7, 1, []string{"D", "j", "D"}},
		{7, 2, concatKeys(
			append([]string{"C"}, append(typeStr("you"), "esc")...),
			[]string{"j", "h", "h", "h", "h"},
			append([]string{"C"}, append(typeStr("know you"), "esc")...),
		)},
		{7, 3, []string{"d", "a", "w", "w", "w", "D"}},
		{7, 4, []string{"d", "a", "w", "j", "b", "D"}},

		{8, 0, []string{"c", "w", "q", "u", "i", "c", "k", "esc", "j", "w", "c", "w", "o", "v", "e", "r", "esc"}},
		{8, 1, []string{"d", "d", "g", "g", "P"}},
		{8, 2, []string{"A", " ", "w", "o", "r", "l", "d", "esc", "G", "o", "S", "e", "e", " ", "y", "o", "u", " ", "s", "o", "o", "n", "esc"}},
		{8, 3, []string{"c", "w", "q", "u", "i", "c", "k", "esc", "w", "w", "w", "w", "c", "w", "o", "v", "e", "r", "esc"}},
		{8, 4, concatKeys(
			[]string{"d", "d", "p"},
			[]string{"j"},
			append([]string{"c", "w"}, append(typeStr("quick"), "esc")...),
			[]string{"j", "w"},
			[]string{"d", "a", "w"},
			[]string{"j", "w", "w"},
			[]string{"D"},
		)},

		{9, 0, []string{"w", "w", "w", "l", "r", "a"}},
		{9, 1, []string{"w", "w", "w", "5", "r", "X"}},
		{9, 2, []string{"w", "l", "l", "l", "~"}},
		{9, 3, []string{"w", "w", "~", "~", "~", "~", "~"}},
		{9, 4, []string{
			"w", "w", "l", "l", "r", "v", "l", "r", "e",
			"j", "b", "~", "~", "~", "~", "~",
			"j", "w", "5", "r", "X",
		}},

		{10, 0, append([]string{"w", "w", "s"}, append(typeStr("five"), "esc")...)},
		{10, 1, append([]string{"w", "w", "w", "3", "s"}, append(typeStr("50"), "esc")...)},
		{10, 2, append([]string{"S"}, append(typeStr("this line is correct now"), "esc")...)},
		{10, 3, append([]string{"2", "S"}, append(typeStr("final combined line"), "esc")...)},
		{10, 4, concatKeys(
			[]string{"w", "w", "w", "2", "s"},
			append(typeStr("one hundred"), "esc"),
			[]string{"j", "0"},
			append([]string{"S"}, append(typeStr("clean sentence here"), "esc")...),
			[]string{"j", "0", "w", "5", "s"},
			append(typeStr("this"), "esc"),
		)},

		{11, 0, concatKeys(
			[]string{"w", "l", "l", "r", "g", "j", "0", "~", "~", "~", "~", "j", "0", "w"},
			append([]string{"s"}, append(typeStr("two"), "esc")...),
		)},
		{11, 1, concatKeys(
			[]string{"d", "d", "p", "j", "2", "S"},
			append(typeStr("final merged line"), "esc"),
		)},
		{11, 2, concatKeys(
			[]string{"w", "w", "5", "r", "X", "j", "0", "S"},
			append(typeStr("clean replacement line"), "esc"),
		)},
		{11, 3, concatKeys(
			[]string{"w", "~", "~", "~", "w", "l", "r", "a", "w", "w"},
			append([]string{"s"}, append(typeStr("four"), "esc")...),
		)},
		{11, 4, []string{
			"d", "d", "p", "j", "0",
			"w", "w", "w", "5", "r", "X",
			"j", "0", "~", "~", "~", "~",
			"j", "0", "w", "w", "w", "D",
		}},

		{12, 0, []string{"e", "X"}},
		{12, 1, []string{"9", "l", "5", "X"}},
		{12, 2, []string{"w", "d", "w", "u", "0", "d", "w"}},
		{12, 3, []string{"d", "w", "d", "w", "d", "w", "u", "u"}},
		{12, 4, []string{
			"4", "l", "X",
			"j", "0",
			"5", "l", "5", "X",
			"j", "0",
			"w", "d", "w", "u", "0", "d", "w",
		}},

		{13, 0, []string{"0"}},
		{13, 1, []string{"^"}},
		{13, 2, []string{"3", "$"}},
		{13, 3, concatKeys(
			[]string{"d", "$", "j", "^", "w", "c", "^"},
			append(typeStr("fix: "), "esc"),
		)},
		{13, 4, concatKeys(
			[]string{"0", "i"},
			append(typeStr("NOTE: "), "esc"),
			[]string{"j", "0", "w", "w", "w", "w", "d", "$"},
			[]string{"j", "^", "w", "c", "^"},
			append(typeStr("DONE: "), "esc"),
		)},

		{14, 0, concatKeys(
			[]string{"$", "b", "c", "w"}, typeStr("target"), []string{"esc"},
			[]string{"j", "d", "d"},
			[]string{"y", "y", "p"},
		)},
		{14, 1, concatKeys(
			[]string{"~", "~", "~", "~"},
			[]string{"j", "0", "w", "w", "c", "i", "w"}, typeStr("spelling"), []string{"esc"},
			[]string{"j", "0", "w", "w", "w", "w", "w", "h", "d", "$"},
			[]string{"j", "0", "s"}, typeStr("three"), []string{"esc"},
			[]string{"j", "0", "w", "3", "r", "X"},
		)},
		{14, 2, concatKeys(
			[]string{"d", "d", "G", "p"},
			[]string{"G", "o"}, typeStr("fourth"), []string{"esc"},
		)},
		{14, 3, concatKeys(
			[]string{"^", "w", "c", "^"}, typeStr("TODO: "), []string{"esc"},
			[]string{"j", "0", "w", "w", "w", "w", "d", "$"},
			[]string{"j", "0", "i"}, typeStr("* "), []string{"esc"},
		)},
		{14, 4, concatKeys(
			[]string{"^", "w", "c", "^"}, typeStr("TODO: "), []string{"esc"},
			[]string{"j", "j", "d", "d"},
			[]string{"w", "w", "c", "i", "w"}, typeStr("typo"), []string{"esc"},
			[]string{"j", "d", "d"},
			[]string{"0", "w", "w", "w", "3", "r", "X"},
			[]string{"j", "0", "~", "~", "~", "~", "~", "~", "~", "~", "~", "~", "~", "~", "~"},
			[]string{"j", "0", "w", "w", "h", "d", "$"},
			[]string{"j", "d", "d", "j", "p", "k", "k", "d", "d", "p"},
			[]string{"j", "j", "o"}, typeStr("closing line added at the end"), []string{"esc"},
		)},

		{15, 0, []string{"4", "w", "2", "b"}},
		{15, 1, []string{"3", "d", "d"}},
		{15, 2, []string{"3", "y", "y", "G", "p"}},
		{15, 3, concatKeys(
			[]string{"d", "3", "w"},
			[]string{"j", "0", "3", "d", "w"},
		)},
		{15, 4, concatKeys(
			[]string{"3", "d", "d"},
			[]string{"j", "0", "d", "3", "w"},
			[]string{"j", "0", "3", "y", "y"},
			[]string{"G", "p"},
		)},

		{16, 0, []string{"d", "i", "("}},
		{16, 1, []string{"d", "a", "("}},
		{16, 2, append([]string{"c", "i", "("}, append(typeStr("42"), "esc")...)},
		{16, 3, append([]string{"c", "i", "\""}, append(typeStr("config.yaml"), "esc")...)},
		{16, 4, concatKeys(
			[]string{"f", "(", "d", "i", "("},
			[]string{"j", "0"},
			[]string{"f", "(", "d", "a", "("},
			[]string{"j", "0"},
			append([]string{"f", "(", "c", "i", "("}, append(typeStr("42"), "esc")...),
			[]string{"j", "0"},
			append([]string{"f", "\"", "c", "i", "\""}, append(typeStr("config.yaml"), "esc")...),
		)},
		{17, 0, append([]string{"c", "i", "["}, append(typeStr("42"), "esc")...)},

		{18, 0, []string{"d", "i", "["}},
		{18, 1, []string{"d", "a", "["}},
		{18, 2, append([]string{"c", "i", "{"}, append(typeStr("debug=true"), "esc")...)},
		{18, 3, append([]string{"c", "i", "'"}, append(typeStr("Devansh"), "esc")...)},
		{18, 4, concatKeys(
			[]string{"f", "[", "d", "i", "["},
			[]string{"j", "0"},
			[]string{"f", "(", "d", "a", "("},
			[]string{"j", "0"},
			append([]string{"f", "(", "c", "i", "("}, append(typeStr("42"), "esc")...),
			[]string{"j", "0"},
			append([]string{"f", "\"", "c", "i", "\""}, append(typeStr("new.txt"), "esc")...),
			[]string{"j", "0"},
			append([]string{"f", "'", "c", "i", "'"}, append(typeStr("Devansh"), "esc")...),
			[]string{"j", "0"},
			append([]string{"f", "{", "c", "i", "{"}, append(typeStr("role=admin"), "esc")...),
		)},

		{19, 0, concatKeys(
			[]string{"y", "i", "w", "$", "a"},
			typeStr(", "),
			[]string{"esc", "p"},
		)},
		{19, 1, []string{"j", "d", "i", "w", "k", "y", "i", "w", "j", "p"}},
		{19, 2, []string{"j", "d", "i", "(", "k", "y", "i", "(", "j", "P"}},
		{19, 3, []string{"j", "d", "i", "\"", "k", "y", "i", "\"", "j", "P"}},
		{19, 4, concatKeys(
			[]string{"j", "0", "w", "w", "d", "i", "w", "k", "0", "w", "w", "y", "i", "w", "j", "0", "w", "w", "P"},
			[]string{"j", "0", "f", "(", "d", "i", "(", "k", "k", "0", "f", "(", "y", "i", "(", "j", "j", "0", "f", "(", "p"},
			[]string{"j", "0", "f", "\"", "d", "i", "\"", "k", "k", "k", "0", "f", "\"", "y", "i", "\"", "j", "j", "j", "0", "f", "\"", "p"},
		)},

		{20, 0, []string{"Y", "p"}},
		{20, 1, []string{"2", "Y", "G", "p"}},
		{20, 2, concatKeys(
			[]string{"Y", "p", "0", "w", "w", "c", "i", "w"},
			typeStr("final"),
			[]string{"esc"},
		)},
		{20, 3, concatKeys(
			[]string{"Y", "p", "0", "f", "(", "c", "i", "("},
			typeStr("5"),
			[]string{"esc"},
		)},
		{20, 4, concatKeys(
			[]string{"3", "Y", "G", "p"},
			[]string{"0", "f", "\"", "c", "i", "\""},
			typeStr("prod"),
			[]string{"esc"},
			[]string{"j", "0", "f", "(", "c", "i", "("},
			typeStr("9"),
			[]string{"esc"},
			[]string{"j", "0", "w", "w", "c", "i", "w"},
			typeStr("admin"),
			[]string{"esc"},
		)},

		{21, 0, concatKeys(
			[]string{"l", "l", "l", "x", "i", "l", "esc", "j", "j", "d", "d", "y", "y", "p", "G", "A"},
			typeStr(" - done"),
			[]string{"esc"},
		)},
		{21, 1, concatKeys(
			[]string{"w", "w", "d", "w"},
			[]string{"j", "0", "w", "c", "w"}, typeStr("new"), []string{"esc"},
			[]string{"j", "0", "w", "c", "i", "w"}, typeStr("outer"), []string{"esc"},
			[]string{"j", "0", "w", "w", "d", "i", "w"},
			[]string{"j", "0", "w", "w", "d", "a", "w"},
			[]string{"j", "0", "w", "w", "e", "l", "D"},
			[]string{"j", "0", "w", "w", "e", "l", "C"}, typeStr(" section"), []string{"esc"},
		)},
		{21, 2, concatKeys(
			[]string{"w", "w", "w", "w", "l", "r", "a"},
			[]string{"j", "0", "~", "~", "~", "~", "~", "w", "~", "~", "~", "~", "~"},
			[]string{"j", "0", "w", "s"}, typeStr("five"), []string{"esc"},
			[]string{"j", "0", "S"}, typeStr("brand new line content"), []string{"esc"},
			[]string{"j", "0", "l", "l", "l", "3", "X"},
			[]string{"j", "0", "w", "d", "w", "u", "0", "d", "w"},
			[]string{"j", "0", "i"}, typeStr("NOTE: "), []string{"esc"},
			[]string{"j", "^", "w", "w", "w", "w", "d", "$"},
			[]string{"j", "^", "w", "c", "^"}, typeStr("DONE: "), []string{"esc"},
		)},
		{21, 3, concatKeys(
			[]string{"3", "d", "d"},
			[]string{"j", "0", "f", "(", "c", "i", "("}, typeStr("newArg"), []string{"esc"},
			[]string{"j", "0", "f", "[", "c", "i", "["}, typeStr("42"), []string{"esc"},
			[]string{"j", "0", "f", "{", "c", "i", "{"}, typeStr("debug=true"), []string{"esc"},
			[]string{"j", "0", "f", "'", "c", "i", "'"}, typeStr("Bob"), []string{"esc"},
			[]string{"j", "0", "w", "w", "j", "d", "i", "w", "k", "y", "i", "w", "j", "p"},
			[]string{"j", "0", "f", "(", "l", "j", "d", "i", "(", "k", "y", "i", "(", "j", "P"},
			[]string{"j", "0", "f", "\"", "l", "j", "d", "i", "\"", "k", "y", "i", "\"", "j", "P"},
			[]string{"j", "0", "Y", "p"},
		)},
		{21, 4, concatKeys(
			[]string{"^", "w", "c", "^"}, typeStr("DONE: "), []string{"esc"},
			[]string{"j", "d", "d"},
			[]string{"j", "0", "w", "c", "i", "w"}, typeStr("found"), []string{"esc"},
			[]string{"j", "d", "d"},
			[]string{"0", "f", "(", "c", "i", "("}, typeStr("XXX"), []string{"esc"},
			[]string{"j", "0", "~", "~", "~", "~", "w", "~", "~", "~", "~", "~", "~", "~", "~"},
			[]string{"j", "0", "w", "w", "e", "l", "D"},
			[]string{"j", "j", "j", "d", "d", "k", "k", "P"},
			[]string{"j", "j", "j", "0", "f", "'", "c", "i", "'"}, typeStr("release"), []string{"esc"},
			[]string{"j", "0", "f", "[", "c", "i", "["}, typeStr("1"), []string{"esc"},
			[]string{"j", "0", "Y", "p", "w", "w", "c", "i", "w"}, typeStr("final"), []string{"esc"},
			[]string{"G", "o"}, typeStr("report complete"), []string{"esc"},
		)},

		{22, 0, []string{"f", "o", "f", "o"}},
		{22, 1, []string{"F", "j", "T", "t"}},
		{22, 2, concatKeys(
			[]string{"c", "t", ":"}, typeStr("note"), []string{"esc"},
			[]string{"j", "0", "d", "f", ":"},
		)},
		{22, 3, concatKeys(
			[]string{"f", ",", ";", "l", "c", "i", "w"}, typeStr("black"), []string{"esc"},
			[]string{"j", "0", "f", ",", ";", ";", ",", "l", "c", "i", "w"}, typeStr("grape"), []string{"esc"},
		)},
		{22, 4, concatKeys(
			[]string{"f", "i", "x", "p"},
			[]string{"j", "0", "f", "P", "c", "t", ":"}, typeStr("DONE"), []string{"esc"},
			[]string{"j", "0", "d", "f", " "},
			[]string{"j", "$", "F", "W", "c", "i", "w"}, typeStr("Error"), []string{"esc"},
			[]string{"j", "0", "f", ":", ";", ";", "l", "c", "i", "w"}, typeStr("90"), []string{"esc"},
		)},

		{23, 0, concatKeys([]string{"/"}, typeStr("ERROR"), []string{"enter"})},
		{23, 1, concatKeys([]string{"?"}, typeStr("configure"), []string{"enter"})},
		{23, 2, concatKeys(
			[]string{"/"}, typeStr("pending"), []string{"enter"},
			[]string{"n", "n", "c", "i", "w"}, typeStr("complete"), []string{"esc"},
		)},
		{23, 3, concatKeys(
			[]string{"?"}, typeStr("ok"), []string{"enter"},
			[]string{"n", "n", "N", "c", "i", "w"}, typeStr("down"), []string{"esc"},
		)},
		{23, 4, concatKeys(
			[]string{"/"}, typeStr("alert"), []string{"enter"},
			[]string{"c", "i", "w"}, typeStr("warning"), []string{"esc"},
			[]string{"n"},
			[]string{"c", "i", "w"}, typeStr("critical"), []string{"esc"},
			[]string{"?"}, typeStr("alert"), []string{"enter"},
			[]string{"c", "i", "w"}, typeStr("resolved"), []string{"esc"},
			[]string{"N"},
			[]string{"c", "i", "w"}, typeStr("monitoring"), []string{"esc"},
		)},

		{24, 0, []string{"%"}},
		{24, 1, []string{"%"}},
		{24, 2, []string{"d", "%"}},
		{24, 3, concatKeys([]string{"c", "%"}, typeStr("REDACTED"), []string{"esc"})},
		{24, 4, concatKeys(
			[]string{"j", "0", "f", "("}, []string{"c", "%"}, typeStr("(45)"), []string{"esc"},
			[]string{"j", "0", "f", "["}, []string{"c", "%"}, typeStr("[5, 10]"), []string{"esc"},
			[]string{"j", "0", "f", "\""}, []string{"c", "i", "\""}, typeStr("staging"), []string{"esc"},
			[]string{"g", "g", "$", "%"}, []string{"o"}, typeStr("# end of settings"), []string{"esc"},
		)},

		{25, 0, []string{"*"}},
		{25, 1, []string{"#", "n"}},
		{25, 2, concatKeys(
			[]string{"*"},
			[]string{"c", "i", "w"}, typeStr("limit"), []string{"esc"},
			[]string{"n"},
			[]string{"c", "i", "w"}, typeStr("limit"), []string{"esc"},
		)},
		{25, 3, concatKeys(
			[]string{"*"},
			[]string{"c", "i", "w"}, typeStr("data"), []string{"esc"},
			[]string{"n"},
			[]string{"c", "i", "w"}, typeStr("data"), []string{"esc"},
			[]string{"n"},
			[]string{"c", "i", "w"}, typeStr("data"), []string{"esc"},
		)},
		{25, 4, concatKeys(
			[]string{"#"},
			[]string{"$", "r", "5"},
			[]string{"j", "0"},
			[]string{"*"},
			[]string{"c", "i", "w"}, typeStr("conn"), []string{"esc"},
			[]string{"n"},
			[]string{"c", "i", "w"}, typeStr("conn"), []string{"esc"},
			[]string{"n"},
			[]string{"c", "i", "w"}, typeStr("conn"), []string{"esc"},
			[]string{"G", "o"}, typeStr("done"), []string{"esc"},
		)},

		{26, 0, concatKeys([]string{"d", "/"}, typeStr("fix"), []string{"enter"})},
		{26, 1, concatKeys([]string{"d", "?"}, typeStr("TODO"), []string{"enter"})},
		{26, 2, concatKeys(
			[]string{"c", "/"}, typeStr("PENDING"), []string{"enter"},
			typeStr("DONE "), []string{"esc"},
		)},
		{26, 3, concatKeys(
			[]string{"y", "/"}, typeStr("method"), []string{"enter"},
			[]string{"j", "P"},
		)},
		{26, 4, concatKeys(
			[]string{"d", "/"}, typeStr("server"), []string{"enter"},
			[]string{"j", "0", "f", "=", "l"},
			[]string{"c", "/"}, typeStr("code"), []string{"enter"},
			typeStr("DONE "), []string{"esc"},
			[]string{"j", "0", "w", "w"},
			[]string{"y", "/"}, typeStr("backoff"), []string{"enter"},
			[]string{"j", "P"},
			[]string{"G", "0", "w", "w", "w"},
			[]string{"d", "?"}, typeStr("TODO"), []string{"enter"},
		)},

		{28, 0, []string{"w", "f", "_", "x", ";", "x"}},
		{28, 1, concatKeys(
			[]string{"/"}, typeStr("TODO"), []string{"enter"},
			[]string{"n", "c", "w"}, typeStr("DONE"), []string{"esc"},
		)},
		{28, 2, concatKeys(
			[]string{"f", "(", "%", "i"}, typeStr(" && y"), []string{"esc"},
		)},
		{28, 3, concatKeys(
			[]string{"*", "c", "i", "w"}, typeStr("cherry"), []string{"esc"},
		)},
		{28, 4, concatKeys(
			[]string{"d", "/"}, typeStr("\\d"), []string{"enter"},
			[]string{"/"}, typeStr("^ERROR"), []string{"enter"},
			[]string{"c", "/"}, typeStr(":"), []string{"enter"}, typeStr("FIXED"), []string{"esc"},
			[]string{"j", "0"},
			[]string{"f", "(", "d", "%"},
		)},

		{29, 0, []string{"v", "l", "l", "l", "d"}},
		{29, 1, []string{"V", "j", "d"}},
		{29, 2, []string{"v", "$", "y", "j", "$", "p"}},
		{29, 3, []string{"v", "e", "~"}},
		{29, 4, concatKeys(
			[]string{"V", "c"}, typeStr("status: done"), []string{"esc"},
		)},
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
			continue
		}
		want := day.Challenges[tc.challenge].Par
		if len(tc.keys) != want {
			t.Errorf("day %d challenge %d (%s): Par=%d but scripted minimal solve took %d keystrokes",
				tc.day, tc.challenge, day.Challenges[tc.challenge].Title, want, len(tc.keys))
		}
	}
}

func concatKeys(groups ...[]string) []string {
	out := []string{}
	for _, g := range groups {
		out = append(out, g...)
	}
	return out
}

func typeStr(s string) []string {
	out := make([]string, 0, len(s))
	for _, r := range s {
		out = append(out, string(r))
	}
	return out
}

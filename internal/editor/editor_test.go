package editor

import "testing"

func feed(b *Buffer, keys ...string) {
	for _, k := range keys {
		b.Input(k)
	}
}

func TestBasicMotion(t *testing.T) {
	b := New([]string{"hello world"}, Pos{0, 0})
	feed(b, "l", "l", "l")
	if b.Cursor.Col != 3 {
		t.Fatalf("expected col 3, got %d", b.Cursor.Col)
	}
	feed(b, "w")
	if b.Cursor.Col != 6 {
		t.Fatalf("expected col 6 after w, got %d", b.Cursor.Col)
	}
	feed(b, "e")
	if b.Cursor.Col != 10 {
		t.Fatalf("expected col 10 after e, got %d", b.Cursor.Col)
	}
	feed(b, "0")
	if b.Cursor.Col != 0 {
		t.Fatalf("expected col 0 after 0, got %d", b.Cursor.Col)
	}
	feed(b, "$")
	if b.Cursor.Col != 10 {
		t.Fatalf("expected col 10 after $, got %d", b.Cursor.Col)
	}
}

func TestDeleteWord(t *testing.T) {
	b := New([]string{"hello world"}, Pos{0, 0})
	feed(b, "d", "w")
	if b.Text() != "world" {
		t.Fatalf("expected 'world', got %q", b.Text())
	}
}

func TestDeleteLineAndPaste(t *testing.T) {
	b := New([]string{"one", "two", "three"}, Pos{0, 0})
	feed(b, "d", "d")
	if b.Text() != "two\nthree" {
		t.Fatalf("expected 'two\\nthree', got %q", b.Text())
	}
	feed(b, "p")
	if b.Text() != "two\none\nthree" {
		t.Fatalf("expected 'two\\none\\nthree', got %q", b.Text())
	}
}

func TestChangeInnerWord(t *testing.T) {
	b := New([]string{"foo (bar) baz"}, Pos{0, 6})
	feed(b, "c", "i", "w")
	if b.Mode != ModeInsert {
		t.Fatalf("expected insert mode after ciw")
	}
	feed(b, "X", "Y", "esc")
	if b.Text() != "foo (XY) baz" {
		t.Fatalf("expected 'foo (XY) baz', got %q", b.Text())
	}
}

func TestChangeInnerParen(t *testing.T) {
	b := New([]string{"foo (bar) baz"}, Pos{0, 6})
	feed(b, "c", "i", "(")
	feed(b, "Z", "esc")
	if b.Text() != "foo (Z) baz" {
		t.Fatalf("expected 'foo (Z) baz', got %q", b.Text())
	}
}

func TestVisualDelete(t *testing.T) {
	b := New([]string{"hello world"}, Pos{0, 0})
	feed(b, "v", "l", "l", "l", "l", "d")
	if b.Text() != " world" {
		t.Fatalf("expected ' world', got %q", b.Text())
	}
}

func TestVisualLineDelete(t *testing.T) {
	b := New([]string{"one", "two", "three"}, Pos{0, 0})
	feed(b, "V", "j", "d")
	if b.Text() != "three" {
		t.Fatalf("expected 'three', got %q", b.Text())
	}
}

func TestUndo(t *testing.T) {
	b := New([]string{"hello world"}, Pos{0, 0})
	feed(b, "d", "w")
	feed(b, "u")
	if b.Text() != "hello world" {
		t.Fatalf("expected undo to restore text, got %q", b.Text())
	}
}

func TestFindChar(t *testing.T) {
	b := New([]string{"hello world"}, Pos{0, 0})
	feed(b, "f", "o")
	if b.Cursor.Col != 4 {
		t.Fatalf("expected col 4 after fo, got %d", b.Cursor.Col)
	}
	// Cursor sits on the 'o' right before the space, so t's target equals
	// the current column; since t is inclusive, dt deletes just that 'o'.
	feed(b, "d", "t", " ")
	if b.Text() != "hell world" {
		t.Fatalf("expected 'hell world' after dt<space>, got %q", b.Text())
	}
}

func TestGGandG(t *testing.T) {
	b := New([]string{"one", "two", "three"}, Pos{2, 0})
	feed(b, "g", "g")
	if b.Cursor.Row != 0 {
		t.Fatalf("expected row 0 after gg, got %d", b.Cursor.Row)
	}
	feed(b, "G")
	if b.Cursor.Row != 2 {
		t.Fatalf("expected row 2 after G, got %d", b.Cursor.Row)
	}
}

func TestCountedMotion(t *testing.T) {
	b := New([]string{"a b c d e"}, Pos{0, 0})
	feed(b, "3", "w")
	if b.Cursor.Col != 6 {
		t.Fatalf("expected col 6 after 3w, got %d", b.Cursor.Col)
	}
}

func TestSubstitute(t *testing.T) {
	b := New([]string{"foo bar foo"}, Pos{0, 0})
	feed(b, ":")
	for _, k := range "s/foo/baz/g" {
		b.Input(string(k))
	}
	feed(b, "enter")
	if b.Text() != "baz bar baz" {
		t.Fatalf("expected 'baz bar baz', got %q", b.Text())
	}
}

func TestSearch(t *testing.T) {
	b := New([]string{"apple banana cherry banana"}, Pos{0, 0})
	feed(b, "/")
	for _, k := range "banana" {
		b.Input(string(k))
	}
	feed(b, "enter")
	if b.Cursor.Col != 6 {
		t.Fatalf("expected col 6 after search, got %d", b.Cursor.Col)
	}
	feed(b, "n")
	if b.Cursor.Col != 20 {
		t.Fatalf("expected col 20 after n, got %d", b.Cursor.Col)
	}
}

func TestMacro(t *testing.T) {
	b := New([]string{"a", "a", "a"}, Pos{0, 0})
	feed(b, "q", "a")
	feed(b, "x", "j")
	feed(b, "q")
	feed(b, "@", "a")
	feed(b, "@", "a")
	if b.Text() != "\n\n" {
		t.Fatalf("expected empty lines after macro replay, got %q", b.Text())
	}
}

func TestPercentMatch(t *testing.T) {
	b := New([]string{"foo(bar(baz))"}, Pos{0, 3})
	feed(b, "%")
	if b.Cursor.Col != 12 {
		t.Fatalf("expected col 12 after %%, got %d", b.Cursor.Col)
	}
}

func TestDollarOperator(t *testing.T) {
	b := New([]string{"hello world"}, Pos{0, 6})
	feed(b, "d", "$")
	if b.Text() != "hello " {
		t.Fatalf("expected 'hello ', got %q", b.Text())
	}
}

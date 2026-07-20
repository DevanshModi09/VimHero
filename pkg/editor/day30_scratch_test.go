package editor

import "testing"

func TestDay30C1(t *testing.T) {
	b := New([]string{"aaTARGETbb"}, Pos{0, 2})
	feed(b, "v", "l", "l", "l", "l", "l", "o", "h", "h", "d")
	t.Logf("keystrokes=%d text=%q", b.Keystrokes, b.Text())
	if b.Text() != "bb" {
		t.Fatalf("expected 'bb', got %q", b.Text())
	}
}

func TestDay30C2(t *testing.T) {
	b := New([]string{"keep1", "cut1", "cut2", "cut3", "keep2"}, Pos{2, 0})
	feed(b, "V", "j", "o", "k", "d")
	t.Logf("keystrokes=%d text=%q", b.Keystrokes, b.Text())
	if b.Text() != "keep1\nkeep2" {
		t.Fatalf("expected 'keep1\\nkeep2', got %q", b.Text())
	}
}

func TestDay30C3(t *testing.T) {
	b := New([]string{"loNOISEhi"}, Pos{0, 2})
	feed(b, "v", "l", "l", "l", "l", "o", "h", "h", "~")
	t.Logf("keystrokes=%d text=%q", b.Keystrokes, b.Text())
	if b.Text() != "LOnoisehi" {
		t.Fatalf("expected 'LOnoisehi', got %q", b.Text())
	}
}

func TestDay30C4(t *testing.T) {
	b := New([]string{"xxTODOxx"}, Pos{0, 2})
	feed(b, "v", "l", "l", "l", "o", "h", "h", "c", "d", "o", "n", "e", "esc")
	t.Logf("keystrokes=%d text=%q mode=%v", b.Keystrokes, b.Text(), b.Mode)
	if b.Text() != "donexx" {
		t.Fatalf("expected 'donexx', got %q", b.Text())
	}
}

func TestDay30C5(t *testing.T) {
	b := New([]string{"aNOTEDb", "copy: "}, Pos{0, 1})
	feed(b, "v", "l", "l", "l", "l", "o", "h", "o", "l", "y", "j", "$", "p")
	t.Logf("keystrokes=%d text=%q", b.Keystrokes, b.Text())
	if b.Text() != "aNOTEDb\ncopy: aNOTEDb" {
		t.Fatalf("expected 'aNOTEDb\\ncopy: aNOTEDb', got %q", b.Text())
	}
}

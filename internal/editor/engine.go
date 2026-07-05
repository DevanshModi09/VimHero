package editor

// Input feeds one key token into the buffer. Key tokens are either a
// single printable rune ("h", "3", "(", ...) or one of the special names:
// "esc", "enter", "backspace", "tab".
func (b *Buffer) Input(key string) {
	if key == "" {
		return
	}
	b.Keystrokes++
	b.KeyLog = append(b.KeyLog, key)
	b.StatusMsg = ""

	if b.recordingReg != 0 && !(b.Mode == ModeNormal && key == "q" && b.pendingKind == 0 && b.pendingOp == 0 && b.count == "") {
		b.macroBuf.WriteString(key + "\x00")
	}

	switch b.Mode {
	case ModeNormal:
		b.handleNormal(key)
	case ModeInsert:
		b.handleInsert(key)
	case ModeVisual, ModeVisualLine:
		b.handleVisual(key)
	case ModeCommand:
		b.handleCommandMode(key)
	}
	b.clampCursor()
}

// key tokens are stored with a trailing NUL as a separator when recording
// macros, since keys can be multi-rune strings like "esc".

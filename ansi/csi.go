package ansi

// See: https://en.wikipedia.org/wiki/ANSI_escape_code#CSI_sequences

var (
	CursorUp       = func(n int) string { return MakeCSISeq("A", n) }
	CursorDown     = func(n int) string { return MakeCSISeq("B", n) }
	CursorForward  = func(n int) string { return MakeCSISeq("C", n) }
	CursorBack     = func(n int) string { return MakeCSISeq("D", n) }
	CursorNextLine = func(n int) string { return MakeCSISeq("E", n) }
	CursorPrevLine = func(n int) string { return MakeCSISeq("F", n) }
	CursorHorzAbs  = func(n int) string { return MakeCSISeq("G", n) }
	CursorPos      = func(n, m int) string { return MakeCSISeq("H", n, m) }
	EraseInDisplay = func(n int) string { return MakeCSISeq("J", n) }
	EraseInLine    = func(n int) string { return MakeCSISeq("K", n) }
	ScrollUp       = func(n int) string { return MakeCSISeq("S", n) }
	ScrollDown     = func(n int) string { return MakeCSISeq("T", n) }
	HorzVertPos    = func(n, m int) string { return MakeCSISeq("f", n, m) }
	SGR            = func(n ...int) string { return MakeCSISeq("m", n...) }
)

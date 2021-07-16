// +build windows

package prompt

import (
	"io"

	colorable "github.com/mattn/go-colorable"
)

// WindowsWriter is a ConsoleWriter implementation for Win32 console.
// Output is converted from VT100 escape sequences by mattn/go-colorable.
type WindowsWriter struct {
	VT100Writer
	out io.Writer
}

// Flush to flush buffer
func (w *WindowsWriter) Flush() error {
	_, err := w.out.Write(w.buffer)
	if err != nil {
		return err
	}
	w.buffer = []byte{}
	return nil
}

var _ ConsoleWriter = &WindowsWriter{}

var (
	// NewStandardOutputWriter is Deprecated: Please use NewStdoutWriter
	NewStandardOutputWriter = NewStdoutWriter
)

// NewStdoutWriter returns ConsoleWriter object to write to stdout.
// This generates win32 control sequences.
func NewStdoutWriter() ConsoleWriter {
	return &WindowsWriter{
		out: colorable.NewColorableStdout(),
	}
}

// NewStderrWriter returns ConsoleWriter object to write to stderr.
// This generates win32 control sequences.
func NewStderrWriter() ConsoleWriter {
	return &WindowsWriter{
		out: colorable.NewColorableStderr(),
	}
}

/* Scrolling */

// ScrollDown scrolls display down one line.
// col is column position of cursor
//
// Windows does not support Index, only Reverse Index
// Index behavior is achieved by writing newline and CUF until at col
// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
func (w *VT100Writer) ScrollDown(col int) {
    c := strconv.Itoa(col)
    w.WriteRaw([]byte{0x0a})
    w.WriteRaw([]byte{0x1b, '['})
    w.WriteRaw([]byte(c))
	w.WriteRaw([]byte{'C'})
}

// ScrollUp scroll display up one line.
func (w *VT100Writer) ScrollUp() {
	w.WriteRaw([]byte{0x1b, 'M'})
}

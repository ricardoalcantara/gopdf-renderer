// git https://github.com/traherom/memstream
// Copied from memstream.go
//
// Package memstream is an expandable ReadWriteSeeker for Golang that works with an
// internally managed byte buffer. Operation is usage is intended to be seamless
// and smooth.
//
// In situations where the maximum read/write sizes are known, a fixed
// []byte/byte buffer will likely offer better performance.
package pdfrenderer

import (
	"errors"
	"io"
)

// MemoryStream is a memory-based, automatically resizing stream that can
// easily fill the role of any file-based IO.
type MemoryStream struct {
	buff []byte
	loc  int
}

// DefaultCapacity is the size in bytes of a new MemoryStream's backing buffer
const DefaultCapacity = 512

// New creates a new MemoryStream instance
func New() *MemoryStream {
	return NewCapacity(DefaultCapacity)
}

// NewCapacity starts the returned MemoryStream with the given capacity
func NewCapacity(cap int) *MemoryStream {
	return &MemoryStream{buff: make([]byte, 0, cap), loc: 0}
}

// Seek sets the offset for the next Read or Write to offset, interpreted
// according to whence: 0 means relative to the origin of the file, 1 means
// relative to the current offset, and 2 means relative to the end. Seek
// returns the new offset and an error, if any.
//
// Seeking to a negative offset is an error. Seeking to any positive offset is
// legal. If the location is beyond the end of the current length, the position
// will be placed at length.
func (m *MemoryStream) Seek(offset int64, whence int) (int64, error) {
	newLoc := m.loc
	switch whence {
	case 0:
		newLoc = int(offset)
	case 1:
		newLoc += int(offset)
	case 2:
		newLoc = len(m.buff) - int(offset)
	}

	if newLoc < 0 {
		return int64(m.loc), errors.New("Unable to seek to a location <0")
	}

	if newLoc > len(m.buff) {
		newLoc = len(m.buff)
	}

	m.loc = newLoc

	return int64(m.loc), nil
}

// Read puts up to len(p) bytes into p. Will return the number of bytes read.
func (m *MemoryStream) Read(p []byte) (n int, err error) {
	n = copy(p, m.buff[m.loc:len(m.buff)])
	m.loc += n

	if m.loc == len(m.buff) {
		return n, io.EOF
	}

	return n, nil
}

// Write writes the given bytes into the memory stream. If needed, the underlying
// buffer will be expanded to fit the new bytes.
func (m *MemoryStream) Write(p []byte) (n int, err error) {
	// Do we have space?
	if available := cap(m.buff) - m.loc; available < len(p) {
		// How much should we expand by?
		addCap := cap(m.buff)
		if addCap < len(p) {
			addCap = len(p)
		}

		newBuff := make([]byte, len(m.buff), cap(m.buff)+addCap)

		copy(newBuff, m.buff)

		m.buff = newBuff
	}

	// Write
	n = copy(m.buff[m.loc:cap(m.buff)], p)
	m.loc += n
	if len(m.buff) < m.loc {
		m.buff = m.buff[:m.loc]
	}

	return n, nil
}

// Bytes returns a copy of ALL valid bytes in the stream, regardless of the current
// position.
func (m *MemoryStream) Bytes() []byte {
	b := make([]byte, len(m.buff))
	copy(b, m.buff)
	return b
}

// Rewind returns the stream to the beginning
func (m *MemoryStream) Rewind() (int64, error) {
	return m.Seek(0, 0)
}

package test

import "io"

type MockConnection struct {
	reader *io.PipeReader
	writer *io.PipeWriter
}

func NewMockConnection() *MockConnection {
	r, w := io.Pipe()
	return &MockConnection{
		reader: r,
		writer: w,
	}
}

func (m *MockConnection) Read(p []byte) (n int, err error) {
	return m.reader.Read(p)
}

func (m *MockConnection) Write(p []byte) (n int, err error) {
	return m.writer.Write(p)
}

func (m *MockConnection) Close() error {
	m.reader.Close()
	m.writer.Close()
	return nil
}

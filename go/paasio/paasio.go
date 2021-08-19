package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	readBytes int64
	readCalls int
	reader    io.Reader
	readMutex sync.Mutex
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.readMutex.Lock()
	defer rc.readMutex.Unlock()

	n, err := rc.reader.Read(p)
	if err == nil {
		rc.readCalls++
		rc.readBytes += int64(n)
	}
	return n, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	rc.readMutex.Lock()
	defer rc.readMutex.Unlock()

	return rc.readBytes, rc.readCalls
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader, readMutex: sync.Mutex{}}
}

type writeCounter struct {
	writtenBytes int64
	writeCalls   int
	writer       io.Writer
	writeMutex   sync.Mutex
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.writeMutex.Lock()
	defer wc.writeMutex.Unlock()

	n, err := wc.writer.Write(p)
	if err == nil {
		wc.writeCalls++
		wc.writtenBytes += int64(n)
	}
	return n, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	wc.writeMutex.Lock()
	defer wc.writeMutex.Unlock()

	return wc.writtenBytes, wc.writeCalls
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{writer: w, writeMutex: sync.Mutex{}}
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter: readCounter{
			reader: rw,
		},
		writeCounter: writeCounter{
			writer: rw,
		},
	}
}

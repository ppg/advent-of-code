package framework

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// readInput will open a streaming io.ReadSeekCloser to read the input file
// according to program flags.
func readInput(w io.Writer) io.ReadSeekCloser {
	flag.Parse()
	input := flag.Arg(0)
	if input == "" {
		input = "input.txt"
	}
	fmt.Fprintf(w, "reading input: %s\n", input)
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return file
}

// readTestInput will read all the test intput into memory so that tests can
// run without file system access.
func readTestInput(w io.Writer) io.ReadSeekCloser {
	reader := readInput(w)
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return &bytesReader{data: data}
}

type bytesReader struct {
	data []byte
	pos  int64
}

func (bytesReader) Close() error { return nil }
func (r *bytesReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.data[r.pos:])
	r.pos += int64(n)
	return n, nil
}

func (r *bytesReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		r.pos = 0 + offset
	case io.SeekCurrent:
		r.pos = r.pos + offset
	case io.SeekEnd:
		r.pos = int64(len(r.data)) + offset
	}
	if r.pos > int64(len(r.data)) {
		return 0, fmt.Errorf("bytesReader: seek out of bounds: %d", r.pos)
	}
	return r.pos, nil
}

func LineParser(line string) string { return strings.TrimSpace(line) }

func ColParser(n int) func(line string) []string {
	return func(line string) []string {
		cols := strings.Split(line, " ")
		if len(cols) != n {
			panic(fmt.Errorf("unexpected number of columns %d, expected %d:\n%s", len(cols), n, line))
		}
		return cols
	}
}

package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestProcessEvents(t *testing.T) {
	old := os.Stdout
	readPipe, writePipe, _ := os.Pipe()
	os.Stdout = writePipe

	ProcessEvents([]Event{
		{4, 456, "Box", 7},
		{5, 466, "AAA", 2},
		{6, 506, "xyz", 3},
	})

	writePipe.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, readPipe)

	expected := "Event 4: average 3, window ends 456\nEvent 5: Invalid checksum\nEvent 6: average 3, window ends 506\n"
	if buf.String() != expected {
		t.Errorf("got: %q but expect: %q", buf.String(), expected)
	}
}

/*

go test -timeout 30s -run ^TestProcessEvents$ github.com/jeamon/goalgo/aggregate-metrics -count=1 -v
=== RUN   TestProcessEvents
--- PASS: TestProcessEvents (0.00s)
PASS
ok      github.com/jeamon/goalgo/aggregate-metrics      0.246s

*/

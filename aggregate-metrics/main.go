package main

import (
	"fmt"
)

type Event struct {
	Identifier int
	Timestamp  int
	Payload    string
	Checksum   int
}

// ProcessEvent processes events and compute avg metric over last 1 min.
// should ignore invalid (wrong checksum) and outside window events.
func ProcessEvents(events []Event) {
	var currentStart int

	for index, event := range events {
		if Checksum(event.Payload) != event.Checksum {
			fmt.Printf("Event %d: Invalid checksum\n", event.Identifier)
			continue
		}

		if event.Timestamp < currentStart {
			fmt.Printf("Event %d: ignored, too late received, window stars %d\n", event.Identifier, currentStart)
			continue
		}

		currentStart = event.Timestamp - 60
		length := 0
		total := 0

		for i := range index + 1 {
			if Checksum(event.Payload) != event.Checksum || events[i].Timestamp < currentStart {
				continue
			}

			total += 1
			length += len(event.Payload)
		}

		avg := length / total
		fmt.Printf("Event %d: average %d, window ends %d\n", event.Identifier, avg, event.Timestamp)
	}
}

func Checksum(payload string) int {
	var sum int
	for _, c := range payload {
		sum += int(c)
	}
	return sum % 10
}

func main() {}

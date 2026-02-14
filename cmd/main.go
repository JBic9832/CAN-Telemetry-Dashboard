package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/jbic9832/can_dashboard/telemetry"
)

func main() {
	rb := telemetry.NewRingBuffer()
	for {
		rb.Produce(rand.IntN(100))
		v := rb.Consume()
		fmt.Printf("Consumed %d\n", v)

		time.Sleep(500 * time.Millisecond)
	}
}

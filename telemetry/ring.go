package telemetry

import "fmt"

type RingBuffer struct {
	consumer int
	producer int
	data     []int
	size     int
	lag      int
}

func NewRingBuffer() *RingBuffer {
	size := 10
	lag := 2
	consumer := 0
	producer := consumer + lag
	data := make([]int, size)
	return &RingBuffer{
		consumer: consumer,
		producer: producer,
		data:     data,
		size:     size,
		lag:      lag,
	}
}

func (r *RingBuffer) Consume() int {
	dataConsumed := r.data[r.consumer%r.size]
	r.consumer++
	return dataConsumed
}

func (r *RingBuffer) Produce(v int) {
	r.data[r.producer%r.size] = v
	r.producer++
	fmt.Printf("Produced %d\n", v)
}

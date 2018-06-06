package heartbeat

import (
	"math/rand"
	"time"
	"log"
	"fmt"
)

// TrueOrFalse generate a bool value based on the chance provided
func TrueOrFalse(chance float64) bool {
	if rand.Float64() < chance {
		return true
	}
	return false
}

// State records the state of the system
type State struct {
	Interval int // unit: millisecond
	Strength int // range 0 - 10
}

// Records is a slice of States
type Records []State

// Deviation introduced to the system
// Deviation defines range of changes of each marker in the State
type Deviation struct {
	Interval int
	Strength int
}

// Probablities define the probilities that the parameters change
type Probablities struct {
	Interval float64
	Strength float64
}

// HBS is the Heart Beat System
type HBS struct {
	expected State
	state State
	deviation Deviation 
	probablities Probablities
	Instruction chan int
	Paused bool
}

// Initialize a HBS system
func Initialize(s State, d Deviation, p Probablities) *HBS {
	h := &HBS{
		expected: s,
		state: s,
		deviation: d,
		probablities: p,
		Instruction: make(chan int, 1),
		Paused: false,
	}
	h.Instruction <- 0
	return h
}

// Begin to run the system
func (h *HBS) Begin() {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		switch <-h.Instruction {
		case 0: // go ahead
			<- time.After(time.Millisecond * time.Duration(h.state.Interval))
			if !h.Paused {
				log.Println(h.state)
				h.Instruction <- 0
			}
			if TrueOrFalse(h.probablities.Interval) {
				if h.state.Interval > h.expected.Interval {
					if TrueOrFalse(0.95) {
						h.state.Interval = h.expected.Interval + rand.Intn(h.deviation.Interval)
					} else {
						h.state.Interval = h.expected.Interval - rand.Intn(h.deviation.Interval)
					}
				} else {
					if TrueOrFalse(0.05) {
						h.state.Interval = h.expected.Interval + rand.Intn(h.deviation.Interval)
					} else {
						h.state.Interval = h.expected.Interval - rand.Intn(h.deviation.Interval)
					}
				}
			}
			if TrueOrFalse(h.probablities.Strength) {
				if h.state.Strength > h.expected.Strength {
					if TrueOrFalse(0.95) {
						h.state.Strength = h.expected.Strength + rand.Intn(h.deviation.Strength)
					} else {
						h.state.Strength = h.expected.Strength - rand.Intn(h.deviation.Strength)
					}
				} else {
					if TrueOrFalse(0.05) {
						h.state.Strength = h.expected.Strength + rand.Intn(h.deviation.Strength)
					} else {
						h.state.Strength = h.expected.Strength - rand.Intn(h.deviation.Strength)
					}
				}
			}
		case 1: // stop for resuming signal
			for {
				c := <-h.Instruction
				if c == 0 {
					h.Paused = false
					fmt.Println("resuming")
					h.Instruction <- 0
					break
				}
			}
		}
	}
}
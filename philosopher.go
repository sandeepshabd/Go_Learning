package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

// Philosopher struct representing each philosopher
type Philosopher struct {
	number    int
	eatCount  int
	leftChop  *sync.Mutex
	rightChop *sync.Mutex
}

// Host struct to manage the number of eating philosophers
type Host struct {
	eatingPhilosophers int
	mutex              sync.Mutex
	permitCh           chan struct{}
}

// Philosopher's eating routine
func (p *Philosopher) dine(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eatCount < 3 {
		// Request permission to eat from the host
		host.requestPermission()

		// Lock both chopsticks (in any order)
		p.leftChop.Lock()
		p.rightChop.Lock()

		// Start eating
		fmt.Printf("starting to eat %d\n", p.number)
		p.eatCount++

		// Simulate eating for a random period
		time.Sleep(time.Second)

		// Finish eating
		fmt.Printf("finishing eating %d\n", p.number)

		// Unlock both chopsticks
		p.rightChop.Unlock()
		p.leftChop.Unlock()

		// Notify host that this philosopher is done eating
		host.doneEating()

		// Simulate thinking for a random period before eating again
		time.Sleep(time.Millisecond * 500)
	}
}

// Host requests permission to eat
func (h *Host) requestPermission() {
	h.permitCh <- struct{}{}
	h.mutex.Lock()
	h.eatingPhilosophers++
	h.mutex.Unlock()
}

// Host is notified when a philosopher finishes eating
func (h *Host) doneEating() {
	h.mutex.Lock()
	h.eatingPhilosophers--
	h.mutex.Unlock()
	<-h.permitCh
}

func main() {
	var wg sync.WaitGroup

	// Initialize chopsticks (mutexes)
	chopsticks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	// Initialize the host, allowing only 2 philosophers to eat concurrently
	host := &Host{
		permitCh: make(chan struct{}, 2), // Buffer size 2 allows up to 2 philosophers to eat concurrently
	}

	// Initialize philosophers
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			number:    i + 1,
			leftChop:  chopsticks[i],
			rightChop: chopsticks[(i+1)%numPhilosophers],
		}
		wg.Add(1)
	}

	// Start philosopher goroutines
	for _, philosopher := range philosophers {
		go philosopher.dine(host, &wg)
	}

	// Wait for all philosophers to finish
	wg.Wait()
}

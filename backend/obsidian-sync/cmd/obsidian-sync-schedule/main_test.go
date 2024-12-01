package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// mockCounter is a thread-safe counter for tracking function calls
type mockCounter struct {
	mu    sync.Mutex
	count int
}

func (m *mockCounter) increment() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count++
}

func (m *mockCounter) getCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.count
}

// counter instance for tests
var counter = &mockCounter{}

// mockApp is a test implementation of AppFunc that uses the thread-safe counter
func mockApp() {
	counter.increment()
}

func TestScheduler(t *testing.T) {
	// Reset counter
	counter = &mockCounter{}

	// Create scheduler with short interval for testing
	scheduler := NewScheduler(100*time.Millisecond, mockApp)

	// Start scheduler in a goroutine
	go scheduler.Start()

	// Let it run for a bit and check if function was called
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, 1, counter.getCount(), "Function should be called once after ticker interval")

	// Stop the scheduler
	scheduler.Stop()
	time.Sleep(50 * time.Millisecond) // Give it time to stop

	// Verify no more calls happened after stopping
	finalCount := counter.getCount()
	time.Sleep(150 * time.Millisecond)
	assert.Equal(t, finalCount, counter.getCount(), "Function should not be called after scheduler is stopped")
}

func TestSchedulerInterval(t *testing.T) {
	// Reset counter
	counter = &mockCounter{}

	// Create scheduler with very short interval for testing
	scheduler := NewScheduler(100*time.Millisecond, mockApp)

	// Start scheduler in a goroutine
	go scheduler.Start()

	// Let it run for enough time to get multiple ticks
	time.Sleep(250 * time.Millisecond)

	// Stop the scheduler
	scheduler.Stop()

	// Verify multiple calls happened at the right interval
	assert.GreaterOrEqual(t, counter.getCount(), 2, "Function should be called multiple times at the configured interval")
}

func TestNewScheduler(t *testing.T) {
	interval := 5 * time.Minute
	scheduler := NewScheduler(interval, mockApp)

	assert.NotNil(t, scheduler, "Scheduler should not be nil")
	assert.Equal(t, interval, scheduler.interval, "Scheduler should have correct interval")
	assert.NotNil(t, scheduler.quit, "Quit channel should not be nil")
	assert.NotNil(t, scheduler.appFunc, "AppFunc should not be nil")
}

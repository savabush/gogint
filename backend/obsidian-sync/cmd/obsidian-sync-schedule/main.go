package main

import (
	"sync"
	"time"

	app "github.com/savabush/obsidian-sync/internal/app"
	. "github.com/savabush/obsidian-sync/internal/config"
)

// AppFunc represents a function that can be scheduled
type AppFunc func()

// Scheduler represents a scheduler that runs a function at regular intervals
type Scheduler struct {
	interval time.Duration
	appFunc  AppFunc
	quit     chan struct{}
	mu       sync.Mutex
	running  bool
}

// NewScheduler creates a new scheduler with the given interval and function
func NewScheduler(interval time.Duration, fn AppFunc) *Scheduler {
	return &Scheduler{
		interval: interval,
		appFunc:  fn,
		quit:     make(chan struct{}),
		running:  false,
	}
}

// Start starts the scheduler
func (s *Scheduler) Start() {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return
	}
	s.running = true
	s.mu.Unlock()

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.appFunc()
		case <-s.quit:
			s.mu.Lock()
			s.running = false
			s.mu.Unlock()
			return
		}
	}
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return
	}
	s.mu.Unlock()
	close(s.quit)
}

// IsRunning returns whether the scheduler is currently running
func (s *Scheduler) IsRunning() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.running
}

// main is the entry point of the obsidian-sync scheduler application.
// It sets up a ticker to run the App() function at regular intervals
// defined by Settings.APP.SCHEDULE. The scheduler can be gracefully
// stopped by sending a signal to the quit channel.
func main() {
	interval := time.Duration(Settings.APP.SCHEDULE) * time.Minute
	Logger.Infof("Starting obsidian-sync scheduler. Starts every %v minutes", Settings.APP.SCHEDULE)
	
	scheduler := NewScheduler(interval, app.App)
	scheduler.Start()
}

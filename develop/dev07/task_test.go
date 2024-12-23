package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	t.Run("single channel", func(t *testing.T) {
		start := time.Now()
		<-or(sig(1 * time.Second))
		if time.Since(start) < 1*time.Second {
			t.Errorf("expected at least 1 second, got %v", time.Since(start))
		}
	})

	t.Run("multiple channels", func(t *testing.T) {
		start := time.Now()
		<-or(sig(2*time.Hour), sig(5*time.Minute), sig(1*time.Second), sig(1*time.Hour), sig(1*time.Minute))
		if time.Since(start) < 1*time.Second {
			t.Errorf("expected at least 1 second, got %v", time.Since(start))
		}
	})
}

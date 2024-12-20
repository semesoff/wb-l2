package main

import (
	"errors"
	"testing"
	"time"

	"github.com/beevik/ntp"
)

type MockTimeProvider struct {
	TimeFunc func(ntpServer string) (time.Time, error)
}

func (m *MockTimeProvider) getTime(ntpServer string) (time.Time, error) {
	return m.TimeFunc(ntpServer)
}

// TestGetTimeSuccess tests the successful case of getTime
func TestGetTimeSuccess(t *testing.T) {
	ntpServer := "time.google.com"

	mockTimeProvider := &MockTimeProvider{
		TimeFunc: ntp.Time,
	}

	currentTime, _ := mockTimeProvider.getTime(ntpServer)

	ntpTime, err := getTime(ntpServer)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !ntpTime.Truncate(time.Second).Equal(currentTime.Truncate(time.Second)) {
		t.Errorf("Expected time %v, got %v", ntpTime, currentTime)
	}
}

func TestGetTimeFailure(t *testing.T) {
	ntpServer := "time.g!oogle.co23!"

	mockTimeProvider := &MockTimeProvider{
		TimeFunc: ntp.Time,
	}

	_, mockErr := mockTimeProvider.getTime(ntpServer)
	_, ntpErr := getTime(ntpServer)

	if errors.Is(ntpErr, mockErr) {
		t.Errorf("Expected error %v, got %v", mockErr, ntpErr)
	}
}

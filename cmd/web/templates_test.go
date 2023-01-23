package main

import (
	"testing"
	"time"

	"github.com/LachlanStephan/ls_server/internal/assert"
)

func TestFormatCreatedAt(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15 (UTC)",
		},
		{
			name: "Unknown",
			tm:   time.Time{},
			want: "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatCreatedAt(tt.tm)

			assert.Equal(t, result, tt.want)
		})
	}
}

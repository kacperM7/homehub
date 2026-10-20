package event

import (
	"errors"
	"testing"
)

func TestParseTopic(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    Topic
		wantErr bool
	}{
		{
			name: "valid home/room/device/metric",
			in:   "home/salon/tv/power",
			want: Topic{room: "salon", device: "tv", metric: "power"},
		},
		{
			name:    "invalid empty",
			in:      "",
			want:    Topic{},
			wantErr: true,
		},
		{
			name:    "invalid wildcard +",
			in:      "home/+/tv/power",
			want:    Topic{},
			wantErr: true,
		},
		{
			name:    "invalid wildcard #",
			in:      "home/room/tv/#",
			want:    Topic{},
			wantErr: true,
		},
		{
			name:    "invalid wrong first segment",
			in:      "garden/salon/tv/power",
			want:    Topic{},
			wantErr: true,
		},
		{
			name:    "invalid empty segment",
			in:      "home//tv/power",
			want:    Topic{},
			wantErr: true,
		},
		{
			name:    "invalid too few segments",
			in:      "home/salon/tv",
			want:    Topic{},
			wantErr: true,
		},
		{
			name:    "invalid too many segments",
			in:      "home/salon/tv/power/extra",
			want:    Topic{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTopic(tt.in)
			if (errors.Is(err, ErrInvalidTopic)) != tt.wantErr {
				t.Fatalf("ParseTopic(%q) error = %v, wantErr %v", tt.in, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseTopic(%q) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

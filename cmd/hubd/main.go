package main

import (
	"log/slog"

	"github.com/kacperM7/homehub/internal/event"
)

func main() {
	slog.Info("Test")

	t := event.Topic{
		Room:   "salon",
		Device: "tv",
		Metric: "power",
	}

	slog.Info("topic", "value", t.String())
}

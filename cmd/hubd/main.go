package main

import (
	"log/slog"

	"github.com/kacperM7/homehub/internal/event"
)

func main() {

	t, err := event.NewTopic("home/salon/tv/power")
	if err != nil {
		slog.Error("topic", "err", err)
		return
	}

	slog.Info("topic", "value", t.String(), "room", t.Room())
}

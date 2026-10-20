package event

import (
	"errors"
	"fmt"
	"regexp"
)

type Topic struct {
	room   string
	device string
	metric string
}

func (t Topic) Room() string   { return t.room }
func (t Topic) Device() string { return t.device }
func (t Topic) Metric() string { return t.metric }

func (t Topic) String() string {
	return fmt.Sprintf("home/%s/%s/%s", t.room, t.device, t.metric)
}

// NewTopic creates new Topic structure from s.
// s must be home/{room}/{device}/{metric}; each segment
// may containg only ASCII letters and digits.
func NewTopic(s string) (Topic, error) {
	t, err := parseTopic(s)
	if err != nil {
		return Topic{}, err
	}
	return t, nil
}

// home / room / device / metric — each segment is letters and digits only.
var topicRe = regexp.MustCompile(`^home/([A-Za-z0-9]+)/([A-Za-z0-9]+)/([A-Za-z0-9]+)$`)

// ErrInvalidTopic is returned when topic string fails the pattern validation
var ErrInvalidTopic = errors.New("invalid topic")

func parseTopic(s string) (Topic, error) {
	m := topicRe.FindStringSubmatch(s)
	if m == nil {
		return Topic{}, fmt.Errorf("parse topic %q: %w", s, ErrInvalidTopic)
	}
	return Topic{
		room:   m[1],
		device: m[2],
		metric: m[3],
	}, nil
}

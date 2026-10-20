package event

import (
	"errors"
	"fmt"
	"regexp"
)

type Topic struct {
	Room   string
	Device string
	Metric string
}

func (t Topic) String() string {
	return fmt.Sprintf("home/%s/%s/%s", t.Room, t.Device, t.Metric)
}

// home / room / device / metric — no MQTT wildcards (+, #) in any section.
var topicRe = regexp.MustCompile(`^home/([^/+#]+)/([^/+#]+)/([^/+#]+)$`)

func ParseTopic(s string) (Topic, error) {
	m := topicRe.FindStringSubmatch(s)
	if m == nil {
		var ErrInvalidTopic = errors.New("invalid topic")
		return Topic{}, fmt.Errorf("topic parsing error %w", ErrInvalidTopic)
	}
	return Topic{
		Room:   m[1],
		Device: m[2],
		Metric: m[3],
	}, nil
}

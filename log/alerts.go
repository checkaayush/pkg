package log

import (
	"fmt"
)

// AlertPriority is a type representing priority of an alert.
type AlertPriority int

const (
	AlertP0 AlertPriority = iota
	AlertP1
	AlertP2
)

// Alert interface is implemented by all alerts.
type Alert interface {
	Priority() AlertPriority
	fmt.Stringer
}

// logAlert implements Alert interface.
type logAlert struct {
	priority      AlertPriority
	servicePrefix string
	alertType     string
}

// Priority returns priority of the alert.
func (l logAlert) Priority() AlertPriority {
	return l.priority
}

// String returns a formatted alert.
func (l logAlert) String() string {
	return fmt.Sprintf("P%d::%s::%s", l.priority, l.servicePrefix, l.alertType)
}

// NewAlert ceates a new instance of logAlert type which implements the
// Alert interface.
func NewAlert(priority AlertPriority, servicePrefix string, alertType string) logAlert {
	return logAlert{
		priority:      priority,
		servicePrefix: servicePrefix,
		alertType:     alertType,
	}
}

func add(a, b int) int {
	return a + b
}

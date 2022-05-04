package notifier

import "strings"

const (
	Info    int = 100
	Success     = 101
	Warning     = 102
	Error       = 103
	Unknown     = 104
)

var typeToIcon = map[int]string{
	Info:    "ℹ️",
	Success: "✅",
	Warning: "⚠️",
	Error:   "❌",
	Unknown: "⁉️",
}

// Notification represents structure of notification
type Notification struct {
	// notificationType represents notification type
	notificationType int

	// notificationTitle represents notification title
	notificationTitle string

	// notificationMessage represents notification message
	notificationMessage string

	// notificationImage represents notification image
	notificationImage string
}

// NotificationOptions represents list of options which can be used to create notification
type NotificationOptions struct {
	// Type represents notification type
	Type int

	// Title represents notification title
	Title string

	// Message represents notification message
	Message string

	// Image represents notification image
	Image string
}

// NewNotification creates new instance of notification
func NewNotification(options NotificationOptions) *Notification {
	instance := &Notification{
		notificationType:    options.Type,
		notificationTitle:   strings.TrimSpace(options.Title),
		notificationMessage: strings.TrimSpace(options.Message),
		notificationImage:   strings.TrimSpace(options.Image),
	}

	if instance.notificationType == 0 {
		instance.notificationType = Unknown
	}

	if instance.notificationTitle == "" {
		instance.notificationTitle = "Notification title"
	}

	if instance.notificationMessage == "" {
		instance.notificationMessage = "Notification Message"
	}

	return instance
}

// Icon returns icon based on provided notification type
func (notification *Notification) Icon() string {
	if value, ok := typeToIcon[notification.notificationType]; ok {
		return value
	}
	return typeToIcon[Unknown]
}

// Type returns notification type
func (notification *Notification) Type() int {
	return notification.notificationType
}

// Title returns notification title
func (notification *Notification) Title() string {
	return notification.notificationTitle
}

// Message returns notification message
func (notification *Notification) Message() string {
	return notification.notificationMessage
}

// Image returns image to be attached to notification
func (notification *Notification) Image() string {
	return notification.notificationImage
}

// HasImage checks whether notification has image set
func (notification *Notification) HasImage() bool {
	return len(notification.Image()) > 0
}

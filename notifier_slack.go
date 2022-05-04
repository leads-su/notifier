package notifier

// slackNotifier defines structure of notifier specific to Telegram
type slackNotifier struct {
	Notifier
}

func NewSlackNotifier() *slackNotifier {
	return &slackNotifier{}
}

// DeliverNotification delivers notification using logic for specific notifier
func (notifier slackNotifier) deliverNotification(notification *Notification) error {

	return nil
}

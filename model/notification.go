package model

type NotificationType int

const (
	NotificationError NotificationType = iota
	NotificationWarning
	NotificationInfo
	NotificationSuccess
)

type Notification struct {
	Model
	Type    NotificationType `json:"type"`
	Title   string           `json:"title"`
	Content string           `json:"content"`
	Details any              `json:"details" gorm:"serializer:json"`
}

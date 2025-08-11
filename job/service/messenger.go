package service

// Messenger is an interface to abstract message sending.
// This methods wrap specific messaging implementations from different libraries.
type Messenger interface {
	GetName() string
	Ping() error
	SendMessage(topic, message string) error
}

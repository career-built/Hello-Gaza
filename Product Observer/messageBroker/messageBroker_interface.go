package messageBroker

// MessageBroker interface
type MessageBroker interface {
	PublishMessages(exchange string, queueName string, messages ...string) error
	ConsumeMessages(queueName string, handler func(message string)) error
	Close() error
}

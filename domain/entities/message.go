package entities

type message struct {
	description string
}

func NewMessage(description string) message {
	m := message{description}

	return m
}

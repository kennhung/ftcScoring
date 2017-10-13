package arena

import (
	"log"
)

// Allow the listeners to buffer a small number of notifications to streamline delivery.
const notifyBufferSize = 3

type Channel struct {
	// The map is essentially a set; the value is ignored.
	listeners map[chan interface{}]struct{}
}

func NewChannel() *Channel {
	Channel := new(Channel)
	Channel.listeners = make(map[chan interface{}]struct{})
	return Channel
}

// Registers and returns a channel that can be read from to receive notification messages. The caller is
// responsible for closing the channel, which will cause it to be reaped from the list of listeners.
func (Channel *Channel) Listen() chan interface{} {
	listener := make(chan interface{}, notifyBufferSize)
	Channel.listeners[listener] = struct{}{}
	return listener
}

// Sends the given message to all registered listeners, and cleans up any listeners that have closed.
func (Channel *Channel) Notify(message interface{}) {
	for listener, _ := range Channel.listeners {
		Channel.notifyListener(listener, message)
	}
}

func (Channel *Channel) notifyListener(listener chan interface{}, message interface{}) {
	defer func() {
		// If channel is closed sending to it will cause a panic; recover and remove it from the list.
		if r := recover(); r != nil {
			delete(Channel.listeners, listener)
		}
	}()

	select {
	case listener <- message:
		// The notification was sent and received successfully.
	default:
		log.Println("Failed to send a notification: blocked listener.")
	}
}

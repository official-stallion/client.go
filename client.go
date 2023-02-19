package go_sdk

import (
	"fmt"
	"net"

	"github.com/official-stallion/go-sdk/internal"
)

// Client is our client handler which enables a user
// to communicate with broker server.
type Client interface {
	// Publish messages to broker
	Publish(topic string, data []byte) error

	// Subscribe over broker
	Subscribe(topic string, handler internal.MessageHandler)

	// Unsubscribe from broker
	Unsubscribe(topic string)
}

// NewClient creates a new client to connect to broker server.
func NewClient(uri string) (Client, error) {
	url, err := urlUnpack(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid uri: %w", err)
	}

	conn, err := net.Dial("tcp", url.address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client, err := internal.NewClient(conn, url.auth)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	return client, nil
}

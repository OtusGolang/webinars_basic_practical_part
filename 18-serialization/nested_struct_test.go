package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Envelope struct {
	From string
	To   string
	Body json.RawMessage
}

type Message struct {
	Subject string
	Content string
}

func TestNestedUnmarshal(t *testing.T) {

	jsonStr := `
{
    "From": "Alice",
    "To": "Bob",
    "Body": {
        "Subject": "Hello",
        "Content": "World"
    }
}`

	envelope := Envelope{}
	err := json.Unmarshal([]byte(jsonStr), &envelope)
	require.Nil(t, err)

	message := Message{}
	err = json.Unmarshal(envelope.Body, &message)
	require.Nil(t, err)

	fmt.Printf(" Envelope: %s\n", envelope)
	fmt.Printf(" Message: %s\n", message)

	// marshal
	message = Message{
		Subject: "Hello",
		Content: "World",
	}
	envelope = Envelope{
		From: "Alice",
		To:   "Bob",
	}
	envelope.Body, err = json.Marshal(message)
	require.Nil(t, err)

	envelopeJson, err := json.Marshal(envelope)
	require.Nil(t, err)

	fmt.Printf(" Envelope: %s\n", envelopeJson)

}

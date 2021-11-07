package main

import (
	"encoding/json"
        
)

type EventHandlerSess func(*Session)
type EventHandler func(*Event)
// EventHandler defines a function which gets passed
// the event as instance pointer

// Event contains events name and data
type Event struct {
	Name string      `json:"event"`
        Data interface{} `json:"data"`
        room string      `json:"room"`
        email string     `json:"email"`
        *Message
         
}

// NewEventFromRaw creates an event object
// from raw binary JSON data
func NewEventFromRaw(rawData []byte) (*Event, error) {
	eve := &Event{}
	err := json.Unmarshal(rawData, eve)
	return eve, err
}

// Raw creates raw binary JSON data from event
// instance
func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}

func (e *Event) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

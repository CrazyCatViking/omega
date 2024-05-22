package gateway

import (
	"encoding/json"

	gatewaymodels "github.com/CrazyCatViking/omega/discord/gateway/models"
)

type EventListener func(payload json.RawMessage)
 
type EventHandler struct {
  eventHandlers map[string][]EventListener
}

func NewEventHandler() *EventHandler {
  return &EventHandler{
    eventHandlers: make(map[string][]EventListener),
  }
}

func (eventHandler *EventHandler) HandleEvent(event gatewaymodels.GatewayEvent) {
  handler, ok := eventHandler.eventHandlers[event.EventName]

  if !ok {
    return
  }

  for _, listener := range handler {
    go listener(event.Data)
  }
}

func (eventHandler *EventHandler) on(event string, listener EventListener) {
  eventHandler.eventHandlers[event] = append(eventHandler.eventHandlers[event], listener)
}

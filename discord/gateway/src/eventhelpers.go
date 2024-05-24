package gateway

import (
	"encoding/json"
	"log"

	gatewaymodels "github.com/CrazyCatViking/omega/discord/gateway/models"
)

func (gateway *DiscordGateway) On(event string, handler func(payload json.RawMessage)) {
  gateway.eventHandler.on(event, handler)
}

func (gateway *DiscordGateway) OnMessageCreate(handler func(event gatewaymodels.MessageCreate)) {
  wrappedHandler := func(payload json.RawMessage) {
    message := gatewaymodels.MessageCreate{}
    err := json.Unmarshal([]byte(payload), &message)
    
    if err != nil {
      log.Println("Error unmarshalling message:", err)
      return
    }

    handler(message)
  }

  gateway.eventHandler.on(MessageCreate, wrappedHandler)
}



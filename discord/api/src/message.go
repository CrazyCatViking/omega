package api

import (
	"log"

	apimodels "github.com/CrazyCatViking/omega/discord/api/models"
)

func (api *DiscordApi) CreateMessage(channelId string, message apimodels.CreateMessageInput) {
  response := api.Post("/channels/" + channelId + "/messages", message)
  
  if (response.Err != nil) {
    log.Println("Error creating message:", response.Err)
    return
  }

  decodedResponse := map[string]interface{}{}
  err := DecodeResponse(response.Response, &decodedResponse)

  if err != nil {
    log.Println("Error decoding response:", err)
    return
  }
}

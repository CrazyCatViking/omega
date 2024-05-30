package api

import (
	"log"

	apimodels "github.com/CrazyCatViking/omega/discord/api/models"
)

func (api *DiscordApi) CreateCommand(applicationId string, commandInput apimodels.CreateCommandInput) {
  response := api.Post("/applications/" + applicationId + "/commands", commandInput)
  
  if (response.Err != nil) {
    log.Println("Error creating command:", response.Err)
    return
  }

  decodedResponse := map[string]interface{}{}
  err := DecodeResponse(response.Response, &decodedResponse)

  if err != nil {
    log.Println("Error decoding response:", err)
    return
  }
}

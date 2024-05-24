package main

import (
	"os"

	apimodels "github.com/CrazyCatViking/omega/discord/api/models"
	api "github.com/CrazyCatViking/omega/discord/api/src"
	gatewaymodels "github.com/CrazyCatViking/omega/discord/gateway/models"
	"github.com/CrazyCatViking/omega/discord/gateway/src"
	dotenv "github.com/CrazyCatViking/omega/utils/dotenv"
)

func main() {
  dotenv.InitEnv()

  token := os.Getenv("BOT_TOKEN")

  gateway := gateway.NewGateway(gateway.ConnectionOptions{
    BotToken: token,
    Intents: 513,
    ShardId: 0,
    ShardCount: 1,
  })

  api := api.NewDiscordApi(token)

  gateway.OnMessageCreate(func(message gatewaymodels.MessageCreate) {
    channelId := message.ChannelId

    if (message.Content == "Hello") {
      api.CreateMessage(channelId, apimodels.CreateMessageInput{
        Content: "World!",
      })
    }
  })

  gateway.Listen()
}

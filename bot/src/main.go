package main

import (
	"log"
	"os"

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

  gateway.OnMessageCreate(func(message gatewaymodels.Message) {
    log.Println("Received message:", message)
  })

  gateway.Listen()
}

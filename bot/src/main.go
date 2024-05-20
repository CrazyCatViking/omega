package main

import (
	"log"
	"os"

	"github.com/CrazyCatViking/omega/discord/gateway/src"
	dotenv "github.com/CrazyCatViking/omega/utils/dotenv"
)

func main() {
  dotenv.InitEnv()

  token := os.Getenv("BOT_TOKEN")
  gatewayUri := os.Getenv("DISCORD_GATEWAY_URI")

  gateway.Test()

  log.Println(token)
  log.Println(gatewayUri)
}

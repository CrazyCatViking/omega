package main

import (
	"log"
	"os"

	dotenv "github.com/CrazyCatViking/omega/utils/src"
)

func main() {
  dotenv.InitEnv()

  token := os.Getenv("BOT_TOKEN")

  log.Println(token)
}

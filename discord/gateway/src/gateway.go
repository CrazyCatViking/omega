package gateway

import (
	"log"
	"os"
	"time"

	gatewaymodels "github.com/CrazyCatViking/omega/discord/gateway/models"
	"github.com/gorilla/websocket"
)

func Test() {
  gatewayUri := os.Getenv("DISCORD_GATEWAY_URI")
  botToken := os.Getenv("BOT_TOKEN")

  log.Println("Connecting to gateway at", gatewayUri)

  connection, _, err := websocket.DefaultDialer.Dial(gatewayUri, nil)

  if err != nil {
    log.Fatal("Error connecting to gateway:", err)
  }

  defer connection.Close()

  handshakeData := gatewaymodels.HandshakeData {
    Token: botToken, 
    Intents: 513,
    Shard: []int{0, 1},
    Properties: gatewaymodels.PropertiesData {
      Os: "linux",
      Browser: "omega",
      Device: "omega",
    },
    Presence: gatewaymodels.PresenceData {
      Status: "online",
      Afk: false,
      Activities: []gatewaymodels.ActivityData {
        {
          Name: "with the API",
          Type: 0,
        },
      },
    },
  }

  payloadData := gatewaymodels.GatewayPayload[gatewaymodels.HandshakeData] {
    OpCode: 2,
    Data: handshakeData,
  }
  
  err = connection.WriteJSON(payloadData)

  go func() {
    for {
      _, message, err := connection.ReadMessage()
      if err != nil {
        log.Println("Error reading message:", err)
        return
      }
      log.Printf("Received: %s\n", message)
    }
  }()

  heartbeatTicker := time.NewTicker(41 * time.Second)
  defer heartbeatTicker.Stop()

  for {
    select {
    case <-heartbeatTicker.C:
      heartbeatPayload := map[string]interface{}{
        "op": 1,
        "d":  nil,
      }
      err = connection.WriteJSON(heartbeatPayload)
      if err != nil {
        log.Println("Error sending heartbeat:", err)
        return
      }
    }
  }
}

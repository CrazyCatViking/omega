package gateway

import (
	"encoding/json"
	"log"
	"os"
	"time"

	gatewaymodels "github.com/CrazyCatViking/omega/discord/gateway/models"
	"github.com/gorilla/websocket"
)

type DiscordGateway struct {
  botToken string
  intents int
  shardId int
  shardCount int

  connection *websocket.Conn

  eventHandler *EventHandler
}

type ConnectionOptions struct {
  BotToken string
  Intents int 
  ShardId int
  ShardCount int
}

func NewGateway(options ConnectionOptions) *DiscordGateway {
  return &DiscordGateway{
    botToken: options.BotToken,
    intents: options.Intents,
    shardId: options.ShardId,
    shardCount: options.ShardCount,
    eventHandler: NewEventHandler(),
  }
}

func (gateway *DiscordGateway) Listen() {
  gatewayUri := os.Getenv("DISCORD_GATEWAY_URI")

  connection, _, err := websocket.DefaultDialer.Dial(gatewayUri, nil)

  if err != nil {
    log.Fatal("Error connecting to gateway:", err)
  }

  gateway.connection = connection
  gateway.performHandshake()

  go gateway.SendHeartbeat();

  for {
    _, message, err := gateway.connection.ReadMessage()

    if err != nil {
      log.Println("Error reading message:", err)
      return
    }

    gatewayEvent := gatewaymodels.GatewayEvent{}
    err = json.Unmarshal(message, &gatewayEvent)

    if err != nil {
      log.Println("Error decoding message:", err)
      return
    }

    switch (gatewayEvent.OpCode) {
      case Dispatch:
        gateway.eventHandler.HandleEvent(gatewayEvent)
      default:
        log.Println("Unhandled opcode:", gatewayEvent.OpCode)
    }
  }
}

func (gateway *DiscordGateway) CloseConnection() {
  gateway.connection.Close()
}

func (gateway *DiscordGateway) performHandshake() error {
  handshakeData := gatewaymodels.HandshakeData {
    Token: gateway.botToken, 
    Intents: gateway.intents,
    Shard: []int{gateway.shardId, gateway.shardCount},
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
          Name: "Initializing...",
          Type: 0,
        },
      },
    },
  }

  payloadData := gatewaymodels.GatewayPayload[gatewaymodels.HandshakeData] {
    OpCode: Identify,
    Data: handshakeData,
  }
  
  err := gateway.connection.WriteJSON(payloadData)

  return err
}

func (gateway *DiscordGateway) SendHeartbeat() {
  heartbeatTicker := time.NewTicker(41 * time.Second)
  defer heartbeatTicker.Stop()

  for {
    select {
    case <-heartbeatTicker.C:
      heartbeatPayload := map[string]interface{}{
        "op": Heartbeat,
        "d":  nil,
      }
      err := gateway.connection.WriteJSON(heartbeatPayload)
      if err != nil {
        log.Println("Error sending heartbeat:", err)
        return
      }
    }
  }
}

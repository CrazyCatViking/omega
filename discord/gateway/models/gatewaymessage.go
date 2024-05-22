package gatewaymodels

import "encoding/json"

type GatewayPayload[T any] struct {
  OpCode int `json:"op"`
  Data T `json:"d"`
}

type GatewayEvent struct {
  OpCode int `json:"op"`
  Data json.RawMessage `json:"d"`
  EventName string `json:"t"`
  SequenceNumber int `json:"s"`
}

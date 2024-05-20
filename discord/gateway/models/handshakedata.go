package gatewaymodels

type HandshakeData struct {
  Token string `json:"token"`
  Properties PropertiesData `json:"properties"`
  Compress bool `json:"compress"`
  LargeThreshold int `json:"large_threshold,omitempty"`
  Shard []int `json:"shard"`
  Presence PresenceData `json:"presence"`
  Intents int `json:"intents"`
}

type PresenceData struct {
  Activities []ActivityData `json:"activities"`
  Status string `json:"status"`
  Since int64 `json:"since,omitempty"`
  Afk bool `json:"afk"`
}

type ActivityData struct {
  Name string `json:"name"`
  Type int `json:"type"`
}

type PropertiesData struct {
  Os string `json:"os"`
  Browser string `json:"browser"`
  Device string `json:"device"`
}

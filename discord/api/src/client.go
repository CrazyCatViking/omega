package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
  Response *http.Response
  Err error
}

type DiscordApi struct {
  botToken string
  client *http.Client
  baseUrl string
}

func NewDiscordApi(token string) *DiscordApi {
  httpClient := &http.Client{}

  return &DiscordApi{
    botToken: token,
    client: httpClient,
    baseUrl: "https://discord.com/api",
  }
}

func (api *DiscordApi) Get(path string, params map[string]string) ApiResponse {
  req, err := http.NewRequest("GET", api.baseUrl + path, nil)

  if err != nil {
    return ApiResponse { Err: err }
  }

  req.Header.Set("Authorization", "Bot " + api.botToken)

  if (params != nil) {
    query := req.URL.Query()

    for key, value := range params {
      query.Add(key, value)
    }

    req.URL.RawQuery = query.Encode()
  }

  res, err := api.client.Do(req)

  if err != nil {
    return ApiResponse { Err: err } 
  }

  return ApiResponse { Response: res }
}

func (api *DiscordApi) Post(path string, body any) ApiResponse {
  marshalledBody, err := json.Marshal(body)

  if err != nil {
    return ApiResponse { Err: err }
  }

  req, err := http.NewRequest("POST", api.baseUrl + path, bytes.NewBuffer(marshalledBody))

  api.setHeaders(req)

  res, err := api.client.Do(req)

  if (err != nil) {
    return ApiResponse{ Err: err }
  }

  return ApiResponse { Response: res } 
}

func (api *DiscordApi) setHeaders(req *http.Request) {
  req.Header.Set("Authorization", "Bot " + api.botToken)
  req.Header.Set("Content-Type", "application/json")
}


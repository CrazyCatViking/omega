package api

import (
	"encoding/json"
	"net/http"
)

func DecodeResponse(res *http.Response, target any) error {
  defer res.Body.Close()
  return json.NewDecoder(res.Body).Decode(target)
}

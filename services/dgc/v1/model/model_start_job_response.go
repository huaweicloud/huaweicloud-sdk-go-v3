package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartJobResponse struct{}"
	}

	return strings.Join([]string{"StartJobResponse", string(data)}, " ")
}

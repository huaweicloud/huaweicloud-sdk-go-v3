package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartTaskResponse struct{}"
	}

	return strings.Join([]string{"StartTaskResponse", string(data)}, " ")
}

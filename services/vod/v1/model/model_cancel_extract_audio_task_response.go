package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CancelExtractAudioTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelExtractAudioTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelExtractAudioTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelExtractAudioTaskResponse", string(data)}, " ")
}

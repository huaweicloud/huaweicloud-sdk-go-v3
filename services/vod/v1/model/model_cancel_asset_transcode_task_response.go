package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CancelAssetTranscodeTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelAssetTranscodeTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelAssetTranscodeTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelAssetTranscodeTaskResponse", string(data)}, " ")
}

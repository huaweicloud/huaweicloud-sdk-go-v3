package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelExtractAudioTaskRequest struct {
	// 媒姿ID

	AssetId string `json:"asset_id"`
}

func (o CancelExtractAudioTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelExtractAudioTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelExtractAudioTaskRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelAssetTranscodeTaskRequest struct {
	// 媒姿ID

	AssetId string `json:"asset_id"`
}

func (o CancelAssetTranscodeTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelAssetTranscodeTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelAssetTranscodeTaskRequest", string(data)}, " ")
}

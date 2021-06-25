package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type PublishAssetFromObsResponse struct {
	// 媒资ID

	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishAssetFromObsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublishAssetFromObsResponse struct{}"
	}

	return strings.Join([]string{"PublishAssetFromObsResponse", string(data)}, " ")
}

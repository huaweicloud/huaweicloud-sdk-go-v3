package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAssetProcessTaskResponse struct {
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetProcessTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetProcessTaskResponse", string(data)}, " ")
}

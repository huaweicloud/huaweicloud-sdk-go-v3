package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAssetsRequest struct {
	// 媒资ID

	AssetId []string `json:"asset_id"`
}

func (o DeleteAssetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAssetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetsRequest", string(data)}, " ")
}

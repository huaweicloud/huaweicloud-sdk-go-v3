package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAssetsRequest struct {
	// 媒资ID，支持一次删除多个媒资，批量删除时以逗号分隔。

	AssetId []string `json:"asset_id"`
}

func (o DeleteAssetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAssetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetsRequest", string(data)}, " ")
}

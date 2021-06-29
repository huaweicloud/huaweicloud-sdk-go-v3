package model

import (
	"encoding/json"

	"strings"
)

type PublishAssetReq struct {
	// 媒资ID。

	AssetId []string `json:"asset_id"`
}

func (o PublishAssetReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublishAssetReq struct{}"
	}

	return strings.Join([]string{"PublishAssetReq", string(data)}, " ")
}

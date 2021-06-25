package model

import (
	"encoding/json"

	"strings"
)

type AssetReviewReq struct {
	// 媒资ID

	AssetId string `json:"asset_id"`

	Review *Review `json:"review"`
}

func (o AssetReviewReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssetReviewReq struct{}"
	}

	return strings.Join([]string{"AssetReviewReq", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAssetListResponse struct {
	// 媒资总数

	Total *int32 `json:"total,omitempty"`
	// 媒资列表

	Assets         *[]AssetSummary `json:"assets,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAssetListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAssetListResponse struct{}"
	}

	return strings.Join([]string{"ListAssetListResponse", string(data)}, " ")
}

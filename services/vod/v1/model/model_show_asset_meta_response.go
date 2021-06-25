package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAssetMetaResponse struct {
	AssetInfoArray *[]AssetInfo `json:"asset_info_array,omitempty"`

	IsTruncated *int32 `json:"is_truncated,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAssetMetaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAssetMetaResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetMetaResponse", string(data)}, " ")
}

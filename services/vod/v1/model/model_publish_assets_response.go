package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type PublishAssetsResponse struct {
	// 发布的媒资信息。

	AssetInfoArray *[]AssetInfo `json:"asset_info_array,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o PublishAssetsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublishAssetsResponse struct{}"
	}

	return strings.Join([]string{"PublishAssetsResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UnpublishAssetsResponse struct {
	AssetInfoArray *[]AssetInfo `json:"asset_info_array,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UnpublishAssetsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnpublishAssetsResponse struct{}"
	}

	return strings.Join([]string{"UnpublishAssetsResponse", string(data)}, " ")
}

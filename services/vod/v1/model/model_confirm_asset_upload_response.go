package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ConfirmAssetUploadResponse struct {
	// 媒资ID

	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmAssetUploadResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmAssetUploadResponse struct{}"
	}

	return strings.Join([]string{"ConfirmAssetUploadResponse", string(data)}, " ")
}

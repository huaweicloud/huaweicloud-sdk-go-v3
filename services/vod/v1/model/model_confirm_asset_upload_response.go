package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmAssetUploadResponse Response Object
type ConfirmAssetUploadResponse struct {

	// 媒资ID
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmAssetUploadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmAssetUploadResponse struct{}"
	}

	return strings.Join([]string{"ConfirmAssetUploadResponse", string(data)}, " ")
}

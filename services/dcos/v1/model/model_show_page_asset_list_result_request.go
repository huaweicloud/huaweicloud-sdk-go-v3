package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPageAssetListResultRequest Request Object
type ShowPageAssetListResultRequest struct {
	Body *AssetListRequestBody `json:"body,omitempty"`
}

func (o ShowPageAssetListResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPageAssetListResultRequest struct{}"
	}

	return strings.Join([]string{"ShowPageAssetListResultRequest", string(data)}, " ")
}

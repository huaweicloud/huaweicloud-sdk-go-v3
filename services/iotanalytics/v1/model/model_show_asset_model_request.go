package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetModelRequest Request Object
type ShowAssetModelRequest struct {

	// 模型ID
	ModelId string `json:"model_id"`
}

func (o ShowAssetModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetModelRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetModelRequest", string(data)}, " ")
}

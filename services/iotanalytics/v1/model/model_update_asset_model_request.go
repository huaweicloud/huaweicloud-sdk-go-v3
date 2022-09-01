package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAssetModelRequest struct {

	// 模型ID
	ModelId string `json:"model_id" xml:"model_id"`

	Body *AssetModelModRequest `json:"body,omitempty" xml:"body"`
}

func (o UpdateAssetModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetModelRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetModelRequest", string(data)}, " ")
}

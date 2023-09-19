package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelInfo 模型信息
type ModelInfo struct {

	// 模型资产ID
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 资产名称
	AssetName *string `json:"asset_name,omitempty"`
}

func (o ModelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelInfo struct{}"
	}

	return strings.Join([]string{"ModelInfo", string(data)}, " ")
}

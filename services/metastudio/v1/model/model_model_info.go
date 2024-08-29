package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelInfo 模型信息
type ModelInfo struct {

	// 模型资产ID，可以从资产库中查询。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 资产名称
	AssetName *string `json:"asset_name,omitempty"`

	// 主播轮换时备选主播数字人资产ID（仅形象资产，不包含音色），可以从资产库中查询。
	BackupModelAssetIds *[]string `json:"backup_model_asset_ids,omitempty"`
}

func (o ModelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelInfo struct{}"
	}

	return strings.Join([]string{"ModelInfo", string(data)}, " ")
}

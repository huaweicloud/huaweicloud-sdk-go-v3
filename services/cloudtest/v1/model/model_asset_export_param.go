package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetExportParam struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`

	// 父节点ID
	ParentId *string `json:"parent_id,omitempty"`

	// 因子列表
	FactorIds *[]string `json:"factor_ids,omitempty"`
}

func (o AssetExportParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExportParam struct{}"
	}

	return strings.Join([]string{"AssetExportParam", string(data)}, " ")
}

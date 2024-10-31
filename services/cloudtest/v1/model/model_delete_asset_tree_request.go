package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetTreeRequest Request Object
type DeleteAssetTreeRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产ID
	AssetId string `json:"asset_id"`

	// 因子目录ID
	ParentId string `json:"parent_id"`
}

func (o DeleteAssetTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetTreeRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetTreeRequest", string(data)}, " ")
}

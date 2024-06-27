package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetTreeRequest Request Object
type ShowAssetTreeRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产库ID
	AssetId string `json:"asset_id"`
}

func (o ShowAssetTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetTreeRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactorByAssetIdRequest Request Object
type ShowFactorByAssetIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产库ID
	AssetId string `json:"asset_id"`

	Body *CommRequestListFactorParam `json:"body,omitempty"`
}

func (o ShowFactorByAssetIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactorByAssetIdRequest struct{}"
	}

	return strings.Join([]string{"ShowFactorByAssetIdRequest", string(data)}, " ")
}

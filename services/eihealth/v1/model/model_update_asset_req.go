package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetReq 更新资产信息请求体
type UpdateAssetReq struct {

	// 短描述
	Summary *string `json:"summary,omitempty"`

	// 长描述
	Description *string `json:"description,omitempty"`

	// 封面图片base64编码
	Picture *string `json:"picture,omitempty"`

	// 标签列表
	Labels *[]string `json:"labels,omitempty"`
}

func (o UpdateAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetReq struct{}"
	}

	return strings.Join([]string{"UpdateAssetReq", string(data)}, " ")
}

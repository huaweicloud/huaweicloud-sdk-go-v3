package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishAssetReq struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 镜像tag
	Tag string `json:"tag"`

	// 资产名称
	Name string `json:"name"`

	// 资产版本
	Version string `json:"version"`

	// 展示名
	Title *string `json:"title,omitempty"`

	// 封面图片base64编码
	Picture *string `json:"picture,omitempty"`

	// 短描述
	Summary *string `json:"summary,omitempty"`

	// 长描述
	Description *string `json:"description,omitempty"`

	// 标签列表
	Labels *[]string `json:"labels,omitempty"`
}

func (o PublishAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAssetReq struct{}"
	}

	return strings.Join([]string{"PublishAssetReq", string(data)}, " ")
}

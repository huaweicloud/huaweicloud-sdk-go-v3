package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishWorkflowReq struct {

	// 资产名称
	Name string `json:"name"`

	// 资产版本
	Version string `json:"version"`

	// 封面图片base64编码
	Picture *string `json:"picture,omitempty"`

	// 展示名
	Title *string `json:"title,omitempty"`

	// 短描述
	Summary *string `json:"summary,omitempty"`

	// 长描述
	Description *string `json:"description,omitempty"`

	// 标签列表
	Labels *[]string `json:"labels,omitempty"`
}

func (o PublishWorkflowReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishWorkflowReq struct{}"
	}

	return strings.Join([]string{"PublishWorkflowReq", string(data)}, " ")
}

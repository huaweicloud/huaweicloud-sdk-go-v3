package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomLlmUpdateBody struct {

	// 自定义大模型名称
	Name *string `json:"name,omitempty"`

	// 自定义大模型描述
	Description *string `json:"description,omitempty"`

	// 模型名称
	ModelName *string `json:"model_name,omitempty"`

	// 协议类型
	Protocol *string `json:"protocol,omitempty"`

	// API地址
	Url *string `json:"url,omitempty"`

	// API密钥
	ApiKey *string `json:"api_key,omitempty"`
}

func (o AiguardCustomLlmUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomLlmUpdateBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomLlmUpdateBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomLlmBody struct {

	// 自定义大模型名称
	Name string `json:"name"`

	// 自定义大模型描述
	Description *string `json:"description,omitempty"`

	// 模型名称
	ModelName string `json:"model_name"`

	// 协议类型
	Protocol string `json:"protocol"`

	// API地址
	Url string `json:"url"`

	// API密钥
	ApiKey string `json:"api_key"`
}

func (o AiguardCustomLlmBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomLlmBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomLlmBody", string(data)}, " ")
}

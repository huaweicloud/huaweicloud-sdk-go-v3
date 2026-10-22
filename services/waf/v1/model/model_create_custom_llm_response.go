package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomLlmResponse Response Object
type CreateCustomLlmResponse struct {

	// 自定义大模型ID
	Id *string `json:"id,omitempty"`

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

	// 引用的策略列表
	ReferencedPolicies *[]ReferencedPolicy `json:"referenced_policies,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateCustomLlmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomLlmResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomLlmResponse", string(data)}, " ")
}

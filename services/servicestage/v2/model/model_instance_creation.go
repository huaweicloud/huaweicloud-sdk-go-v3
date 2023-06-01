package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建实例请求体
type InstanceCreation struct {

	// 实例名称，必填
	Name *string `json:"name,omitempty"`

	// 实例描述，非必填
	Desc *string `json:"desc,omitempty"`

	// 架构部署图，通过设计器创建时必填
	Diagram *string `json:"diagram,omitempty"`

	// 设计器架构图base64编码,通过设计器创建或更新实例时可选
	Image *string `json:"image,omitempty"`

	// 预置模板id，通过模板创建时必填
	TemplateId *string `json:"template_id,omitempty"`

	Variables *InstanceCreationVariables `json:"variables,omitempty"`

	// 实例id,更新时必填
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o InstanceCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceCreation struct{}"
	}

	return strings.Join([]string{"InstanceCreation", string(data)}, " ")
}

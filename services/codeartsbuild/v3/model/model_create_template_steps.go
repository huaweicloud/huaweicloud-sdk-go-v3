package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateSteps 构建执行的步骤，采用驼峰式命名
type CreateTemplateSteps struct {

	// 具体的构建步骤
	Properties map[string]interface{} `json:"properties,omitempty"`

	// 构建模块id
	ModuleId string `json:"module_id"`

	// 构建模块名称
	Name *string `json:"name,omitempty"`

	// 构建版本
	Version *string `json:"version,omitempty"`

	// 是否开启
	Enable *bool `json:"enable,omitempty"`
}

func (o CreateTemplateSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateSteps struct{}"
	}

	return strings.Join([]string{"CreateTemplateSteps", string(data)}, " ")
}

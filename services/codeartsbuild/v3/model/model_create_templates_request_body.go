package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplatesRequestBody 创建作业模板接口请求体
type CreateTemplatesRequestBody struct {
	Template *CreateTemplate `json:"template"`

	// 模板命名
	Name string `json:"name"`

	// 模板说明
	Description *string `json:"description,omitempty"`

	// 工具类型
	ToolType *string `json:"tool_type,omitempty"`

	// 构建执行参数列表
	Parameters *[]CreateBuildJobParameter `json:"parameters,omitempty"`
}

func (o CreateTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemplatesRequestBody", string(data)}, " ")
}

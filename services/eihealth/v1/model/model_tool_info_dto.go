package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ToolInfoDto struct {

	// 作业依赖的组件id
	ToolId *string `json:"tool_id,omitempty"`

	// 作业依赖的组件名称
	ToolName *string `json:"tool_name,omitempty"`

	// 作业依赖的组件版本
	ToolVersion *string `json:"tool_version,omitempty"`

	// 作业依赖的组件类型，取值范围app|workflow
	ToolType *string `json:"tool_type,omitempty"`
}

func (o ToolInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ToolInfoDto struct{}"
	}

	return strings.Join([]string{"ToolInfoDto", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowTemplateConfigsVo 状态流流转线流转配置信息
type WorkflowTemplateConfigsVo struct {

	// 操作项配置
	ConfigValue *[]map[string]interface{} `json:"configValue,omitempty"`
}

func (o WorkflowTemplateConfigsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowTemplateConfigsVo struct{}"
	}

	return strings.Join([]string{"WorkflowTemplateConfigsVo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowTemplateNodesVo 状态流中状态信息
type WorkflowTemplateNodesVo struct {

	// 状态编码
	Code *string `json:"code,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`
}

func (o WorkflowTemplateNodesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowTemplateNodesVo struct{}"
	}

	return strings.Join([]string{"WorkflowTemplateNodesVo", string(data)}, " ")
}

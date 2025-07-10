package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowTemplateVo 工作流信息
type WorkflowTemplateVo struct {

	// 状态流中的状态信息
	ProcessNodes *[]WorkflowTemplateNodesVo `json:"processNodes,omitempty"`

	// 状态流中的流转线信息
	ProcessFlows *[]WorkflowTemplateFlowsVo `json:"processFlows,omitempty"`
}

func (o WorkflowTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowTemplateVo struct{}"
	}

	return strings.Join([]string{"WorkflowTemplateVo", string(data)}, " ")
}

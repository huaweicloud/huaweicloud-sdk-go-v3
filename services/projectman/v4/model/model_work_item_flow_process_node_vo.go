package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowProcessNodeVo 工作流节点信息
type WorkItemFlowProcessNodeVo struct {

	// 节点ID
	Id *string `json:"id,omitempty"`

	// 节点类别
	Category *string `json:"category,omitempty"`

	// 工作流实例ID
	ProcessInstanceId *string `json:"process_instance_id,omitempty"`

	// 工作流活动ID
	WorkflowActivityId *string `json:"workflow_activity_id,omitempty"`

	// 节点编码
	Code *string `json:"code,omitempty"`

	Config *WorkItemFlowNodeConfigVo `json:"config,omitempty"`

	// 是否允许挂起
	EnableSuspend *bool `json:"enable_suspend,omitempty"`
}

func (o WorkItemFlowProcessNodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowProcessNodeVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowProcessNodeVo", string(data)}, " ")
}

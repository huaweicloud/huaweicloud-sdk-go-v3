package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowInstanceTopology 流程实例的拓扑图
type WorkflowInstanceTopology struct {

	// **参数解释**: 拓扑图节点信息 **约束限制**: 不涉及
	AuditLogs *[]TopologyNodeInfo `json:"audit_logs,omitempty"`
}

func (o WorkflowInstanceTopology) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowInstanceTopology struct{}"
	}

	return strings.Join([]string{"WorkflowInstanceTopology", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbsFlowEntry struct {

	// 问题流程id
	QaFlowId *string `json:"qa_flow_id,omitempty"`

	// 开始节点id
	StartNodeId *string `json:"start_node_id,omitempty"`
}

func (o CbsFlowEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbsFlowEntry struct{}"
	}

	return strings.Join([]string{"CbsFlowEntry", string(data)}, " ")
}

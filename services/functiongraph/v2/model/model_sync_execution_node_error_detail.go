package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncExecutionNodeErrorDetail 函数流节点错误信息
type SyncExecutionNodeErrorDetail struct {

	// 流程节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 错误详细信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 流程实例创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o SyncExecutionNodeErrorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncExecutionNodeErrorDetail struct{}"
	}

	return strings.Join([]string{"SyncExecutionNodeErrorDetail", string(data)}, " ")
}

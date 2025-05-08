package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionInstancesRequest Request Object
type ListExecutionInstancesRequest struct {

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 从offset指定的下一条数据开始查询
	Offset *int32 `json:"offset,omitempty"`

	// 工单步骤id
	ExecutionStepId *string `json:"execution_step_id,omitempty"`
}

func (o ListExecutionInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListExecutionInstancesRequest", string(data)}, " ")
}

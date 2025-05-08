package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionStepsRequest Request Object
type ListExecutionStepsRequest struct {
	ExecutionId string `json:"execution_id"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 从offset指定的下一条数据开始查询
	Offset *int64 `json:"offset,omitempty"`

	// 步骤id数组
	ExecutionStepIdList *[]string `json:"execution_step_id_list,omitempty"`
}

func (o ListExecutionStepsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionStepsRequest struct{}"
	}

	return strings.Join([]string{"ListExecutionStepsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateExecutionRequestBody struct {

	// 工单id
	ExecutionId string `json:"execution_id"`

	// 操作类型，枚举项RESUME,TERMINATE,RETRY,SKIP_BATCH
	OperateType string `json:"operate_type"`
}

func (o OperateExecutionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateExecutionRequestBody struct{}"
	}

	return strings.Join([]string{"OperateExecutionRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateExecutionRequest Request Object
type OperateExecutionRequest struct {
	Body *OperateExecutionRequestBody `json:"body,omitempty"`
}

func (o OperateExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateExecutionRequest struct{}"
	}

	return strings.Join([]string{"OperateExecutionRequest", string(data)}, " ")
}

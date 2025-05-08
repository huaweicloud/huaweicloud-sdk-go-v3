package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateExecutionResponse Response Object
type OperateExecutionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o OperateExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateExecutionResponse struct{}"
	}

	return strings.Join([]string{"OperateExecutionResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionStepInputs struct {

	// 参数的key
	Key *string `json:"key,omitempty"`

	// 参数的value
	Value *string `json:"value,omitempty"`
}

func (o ExecutionStepInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionStepInputs struct{}"
	}

	return strings.Join([]string{"ExecutionStepInputs", string(data)}, " ")
}

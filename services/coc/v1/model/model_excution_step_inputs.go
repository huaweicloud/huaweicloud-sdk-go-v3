package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExcutionStepInputs struct {

	// 参数的key
	Key *string `json:"key,omitempty"`

	// 参数的value
	Value *string `json:"value,omitempty"`
}

func (o ExcutionStepInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExcutionStepInputs struct{}"
	}

	return strings.Join([]string{"ExcutionStepInputs", string(data)}, " ")
}

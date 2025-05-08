package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListExecutionResponseParameters struct {

	// 参数名
	Key *string `json:"key,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`
}

func (o ListExecutionResponseParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionResponseParameters struct{}"
	}

	return strings.Join([]string{"ListExecutionResponseParameters", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyParameterValue 规则参数值
type PolicyParameterValue struct {

	// 规则参数值
	Value *interface{} `json:"value,omitempty"`
}

func (o PolicyParameterValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyParameterValue struct{}"
	}

	return strings.Join([]string{"PolicyParameterValue", string(data)}, " ")
}

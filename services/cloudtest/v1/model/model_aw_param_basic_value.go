package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AwParamBasicValue struct {

	// 参数默认值，用例有效
	Value *string `json:"value,omitempty"`

	// 参数值范围，逻辑用例有效
	ValueRange *string `json:"value_range,omitempty"`
}

func (o AwParamBasicValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwParamBasicValue struct{}"
	}

	return strings.Join([]string{"AwParamBasicValue", string(data)}, " ")
}

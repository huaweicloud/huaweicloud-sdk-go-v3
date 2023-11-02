package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataobjectSearchConditionData struct {

	// 字段
	Filed *string `json:"filed,omitempty"`

	// 逻辑表达式
	Expression *string `json:"expression,omitempty"`

	// 字段值
	Value *string `json:"value,omitempty"`
}

func (o DataobjectSearchConditionData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectSearchConditionData struct{}"
	}

	return strings.Join([]string{"DataobjectSearchConditionData", string(data)}, " ")
}

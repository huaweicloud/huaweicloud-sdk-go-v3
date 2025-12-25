package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaselineSearchRequestBodyConditionConditions struct {

	// 表达式名称
	Name *string `json:"name,omitempty"`

	// 表达式内容列表
	Data *[]string `json:"data,omitempty"`
}

func (o BaselineSearchRequestBodyConditionConditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineSearchRequestBodyConditionConditions struct{}"
	}

	return strings.Join([]string{"BaselineSearchRequestBodyConditionConditions", string(data)}, " ")
}

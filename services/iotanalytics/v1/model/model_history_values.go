package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryValues struct {
	// 属性名称

	PropertyName *string `json:"property_name,omitempty"`
	// 属性值

	Values *[]interface{} `json:"values,omitempty"`
}

func (o HistoryValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryValues struct{}"
	}

	return strings.Join([]string{"HistoryValues", string(data)}, " ")
}

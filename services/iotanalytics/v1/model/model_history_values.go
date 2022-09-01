package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryValues struct {

	// 属性名称
	PropertyName *string `json:"property_name,omitempty" xml:"property_name"`

	// 属性值
	Values *[]interface{} `json:"values,omitempty" xml:"values"`
}

func (o HistoryValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryValues struct{}"
	}

	return strings.Join([]string{"HistoryValues", string(data)}, " ")
}

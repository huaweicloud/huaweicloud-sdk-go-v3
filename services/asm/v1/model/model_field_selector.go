package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldSelector “字段选择器（Field selectors）”允许你根据一个或多个资源字段的值筛选 Kubernetes 对象
type FieldSelector struct {

	// Key值
	Key *string `json:"key,omitempty"`

	// 操作符，仅支持取值\"In\"
	Operator *string `json:"operator,omitempty"`

	// Value值
	Values *[]string `json:"values,omitempty"`
}

func (o FieldSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldSelector struct{}"
	}

	return strings.Join([]string{"FieldSelector", string(data)}, " ")
}

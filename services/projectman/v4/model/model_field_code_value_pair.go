package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldCodeValuePair 工作项自定义字段的code-value映射
type FieldCodeValuePair struct {

	// 工作项字段code值
	Code *string `json:"code,omitempty"`

	// 工作项自定义字段值，为多选时用英文逗号分隔
	Value *string `json:"value,omitempty"`
}

func (o FieldCodeValuePair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldCodeValuePair struct{}"
	}

	return strings.Join([]string{"FieldCodeValuePair", string(data)}, " ")
}

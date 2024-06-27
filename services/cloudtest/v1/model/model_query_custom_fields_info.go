package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryCustomFieldsInfo 用例自定义字段信息
type QueryCustomFieldsInfo struct {

	// 测试用例自定义字段Id
	Id *int32 `json:"id,omitempty"`

	// 测试用例自定义字段值
	Values *[]string `json:"values,omitempty"`

	// 自定义字段名，优先取id再取fieldName
	FieldName *string `json:"field_name,omitempty"`
}

func (o QueryCustomFieldsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCustomFieldsInfo struct{}"
	}

	return strings.Join([]string{"QueryCustomFieldsInfo", string(data)}, " ")
}

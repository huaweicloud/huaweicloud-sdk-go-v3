package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbParameter 数据库配置参数。
type DbParameter struct {

	// 参数名称。
	ParamName *string `json:"param_name,omitempty"`

	// 参数类型。
	DataType *string `json:"data_type,omitempty"`

	// 参数默认值。
	DefaultValue *string `json:"default_value,omitempty"`

	// 参数取值范围。
	ValueRange *string `json:"value_range,omitempty"`

	// 参数描述。
	Description *string `json:"description,omitempty"`
}

func (o DbParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbParameter struct{}"
	}

	return strings.Join([]string{"DbParameter", string(data)}, " ")
}

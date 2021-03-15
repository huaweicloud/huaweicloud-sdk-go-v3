package model

import (
	"encoding/json"

	"strings"
)

// 参数服务对象。
type ServiceCommandPara struct {
	// 参数的名称。
	ParaName string `json:"para_name"`
	// 参数的数据类型。取值范围：int，long，decimal，string，DateTime，jsonObject，enum，boolean，string list。
	DataType string `json:"data_type"`
	// 参数是否必选。默认为false。
	Required *bool `json:"required,omitempty"`
	// 参数的枚举值列表。
	EnumList *[]string `json:"enum_list,omitempty"`
	// 参数的最小值。
	Min *string `json:"min,omitempty"`
	// 参数的最大值。
	Max *string `json:"max,omitempty"`
	// 参数的最大长度。
	MaxLength *int32 `json:"max_length,omitempty"`
	// 参数的步长。
	Step *float64 `json:"step,omitempty"`
	// 参数的单位。
	Unit *string `json:"unit,omitempty"`
	// 参数的描述。
	Description *string `json:"description,omitempty"`
}

func (o ServiceCommandPara) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceCommandPara struct{}"
	}

	return strings.Join([]string{"ServiceCommandPara", string(data)}, " ")
}

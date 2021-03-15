package model

import (
	"encoding/json"

	"strings"
)

// 属性服务对象。
type ServiceProperty struct {
	// 设备属性名称。
	PropertyName string `json:"property_name"`
	// 设备属性的数据类型。取值范围：int，long，decimal，string，DateTime，jsonObject，enum，boolean，string list。
	DataType string `json:"data_type"`
	// 设备属性是否必选。默认为false。
	Required *bool `json:"required,omitempty"`
	// 设备属性的枚举值列表。
	EnumList *[]string `json:"enum_list,omitempty"`
	// 设备属性的最小值。
	Min *string `json:"min,omitempty"`
	// 设备属性的最大值。
	Max *string `json:"max,omitempty"`
	// 设备属性的最大长度。
	MaxLength *int32 `json:"max_length,omitempty"`
	// 设备属性的步长。
	Step *float64 `json:"step,omitempty"`
	// 设备属性的单位。
	Unit *string `json:"unit,omitempty"`
	// 设备属性的访问模式。取值范围：RWE，RW，RE，WE，E，W，R。 - R：属性值可读 - W：属性值可写 - E：属性值可订阅，即属性值变化时上报事件
	Method string `json:"method"`
	// 设备属性的描述。
	Description *string `json:"description,omitempty"`
	// 设备属性的默认值。如果设置了默认值，使用该产品创建设备时，会将该属性的默认值写入到该设备的设备影子预期数据中，待设备上线时将该属性默认值下发给设备。
	DefaultValue *interface{} `json:"default_value,omitempty"`
}

func (o ServiceProperty) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceProperty struct{}"
	}

	return strings.Join([]string{"ServiceProperty", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArrowType Arrow的类型定义，支持简单类型和复杂类型，遵循Apache Arrow类型系统标准。
type ArrowType struct {

	// Arrow类型名称，如Int, Float, Utf8, Bool, Date, Timestamp, List, Struct等。
	Name string `json:"name"`

	// 类型的位宽，适用于Int、Float等数值类型。如Int32的bit_width为32。
	BitWidth *int32 `json:"bit_width,omitempty"`

	// 是否为有符号类型，适用于整数类型。
	IsSigned *bool `json:"is_signed,omitempty"`

	// Decimal类型的精度。
	Precision *int32 `json:"precision,omitempty"`

	// Decimal类型的标度。
	Scale *int32 `json:"scale,omitempty"`

	// 时间单位，适用于Date、Timestamp、Time类型。如SECOND、MILLISECOND、MICROSECOND、NANOSECOND。
	Unit *string `json:"unit,omitempty"`

	// 时区信息，适用于Timestamp类型。
	Timezone *string `json:"timezone,omitempty"`

	// FixedSizeList类型的列表大小，表示固定大小列表中的元素数量。
	ListSize *int32 `json:"list_size,omitempty"`
}

func (o ArrowType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArrowType struct{}"
	}

	return strings.Join([]string{"ArrowType", string(data)}, " ")
}

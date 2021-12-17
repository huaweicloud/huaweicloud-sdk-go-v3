package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据格式
type DataSchema struct {
	// integer或double类型的的最小值，当属性值超过范围时系统不予存储，integer时范围：-9223372036854775808 ~ 9223372036854775807；double时范围：-1.7976931348623157E308 ~ 1.7976931348623157E308

	Min float32 `json:"min,omitempty"`
	// integer或double类型的最大值，当属性值超过范围时系统不予存储，integer时范围：-9223372036854775808 ~ 9223372036854775807；double时范围：-1.7976931348623157E308 ~ 1.7976931348623157E308

	Max float32 `json:"max,omitempty"`
	// string类型的最小长度，当属性值超过范围时系统不予存储，范围：0 ~ 4096

	MinLength *int32 `json:"min_length,omitempty"`
	// string类型的最大长度，当属性值超过范围时系统不予存储，范围：0 ~ 4096

	MaxLength *int32 `json:"max_length,omitempty"`
	// 数据类型，字符串（string）、整数（integer）、浮点数（double）、对象（object），系统支持的对象型存储大小为1MB，超过时不予存储

	DataType string `json:"data_type"`
}

func (o DataSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSchema struct{}"
	}

	return strings.Join([]string{"DataSchema", string(data)}, " ")
}

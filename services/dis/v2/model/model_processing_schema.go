package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessingSchema 根据源数据的时间戳和已配置的\"partition_format\"生成对应的转储时间目录。将源数据的时间戳使用“yyyy/MM/dd/HH/mm”格式生成分区字符串，用来定义写到OBS的Object文件所在的目录层次结构。
type ProcessingSchema struct {

	// 源数据时间戳的属性名称。
	TimestampName string `json:"timestamp_name"`

	// 源数据时间戳的类型。  - String - Timestamp：Long类型的13位时间戳
	TimestampType string `json:"timestamp_type"`

	// 源数据时间戳的类型为String时必选，用于根据时间戳格式生成OBS的时间目录。
	TimestampFormat *string `json:"timestamp_format,omitempty"`
}

func (o ProcessingSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessingSchema struct{}"
	}

	return strings.Join([]string{"ProcessingSchema", string(data)}, " ")
}

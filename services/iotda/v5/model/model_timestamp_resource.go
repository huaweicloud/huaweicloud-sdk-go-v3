package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimestampResource 是否校验设备时间戳，若不为空则表示校验，如果设备连接参数（clientId、username）中包含时间戳建议开启校验。开启校验平台会对比设备携带时间戳与平台系统时间，若设备时间戳加一小时小于平台系统时间则校验失败。若想关闭校验则把value属性的值设置为空json:{}
type TimestampResource struct {

	// UNIX：表示为Unix时间戳秒级别长整数，FORMAT：表示为特定时间格式，需要在format字段中指定具体格式。
	Type *string `json:"type,omitempty"`

	// 时间日期格式。
	Pattern *string `json:"pattern,omitempty"`

	// 设备连接时具体时间戳值，可使用FunctionDefinition下的函数编程时间戳的取值，若想关闭时间戳校验则该字段值设置为空json:{}。
	Value *interface{} `json:"value,omitempty"`
}

func (o TimestampResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimestampResource struct{}"
	}

	return strings.Join([]string{"TimestampResource", string(data)}, " ")
}

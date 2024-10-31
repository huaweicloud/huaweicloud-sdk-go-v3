package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Processor struct {

	// 解析器类型 processor_regex 正则解析 processor_split_string 分词符 processor_json json解析器类型 processor_gotime自定义时间类型 processor_filter_regex日志过滤 processor_drop删除字段类型 processor_rename修改字段类型
	Type *string `json:"type,omitempty"`

	Detail *Detail `json:"detail,omitempty"`
}

func (o Processor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Processor struct{}"
	}

	return strings.Join([]string{"Processor", string(data)}, " ")
}

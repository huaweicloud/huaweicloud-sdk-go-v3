package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsCommonConfig OBS通用配置
type ObsCommonConfig struct {

	// Map<String, String>结构
	ColumnMap *interface{} `json:"column_map,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 分隔符
	Delimiter *string `json:"delimiter,omitempty"`

	// 引用
	Quote *string `json:"quote,omitempty"`

	// 规避
	Escape *string `json:"escape,omitempty"`

	// 是否是标头
	Header *bool `json:"header,omitempty"`

	// 数据类型
	DataType *string `json:"data_type,omitempty"`

	// 数据格式
	DateFormat *string `json:"date_format,omitempty"`

	// 时间格式
	TimestampFormat *string `json:"timestamp_format,omitempty"`

	// 为空时默认值
	NullValue *string `json:"null_value,omitempty"`

	// 注解
	Comment *string `json:"comment,omitempty"`

	// 解析模式
	ParseMode *string `json:"parse_mode,omitempty"`

	// 联表
	JoinTable *string `json:"join_table,omitempty"`
}

func (o ObsCommonConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsCommonConfig struct{}"
	}

	return strings.Join([]string{"ObsCommonConfig", string(data)}, " ")
}

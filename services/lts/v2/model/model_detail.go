package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Detail struct {

	// 自定义时间key字段名称
	SourceKey *string `json:"source_key,omitempty"`

	// 正则解析正则表达式 单行完全正则，多行完全正则需填此字段
	Regex *string `json:"regex,omitempty"`

	// 字段名称列表 与field_name字段保持一致
	Keys *[]string `json:"keys,omitempty"`

	// 首行正则表达式
	MultiLineRegex *string `json:"multi_line_regex,omitempty"`

	// 是否上传原始日志
	KeepSource *bool `json:"keep_source,omitempty"`

	// 是否上传解析失败日志
	KeepSourceIfParseError *bool `json:"keep_source_if_parse_error,omitempty"`

	// 分隔符 自定义字符最大长度10，自定义字符串最大长度30
	SplitSep *string `json:"split_sep,omitempty"`

	// 分隔符类型：char-自定义字符；special_char-不可见字符；string-自定义字符串
	SplitType *string `json:"split_type,omitempty"`

	// 添加的字段列表：<字段名称 : 字段内容>
	Fields map[string]string `json:"fields,omitempty"`

	// 删除的字段列表
	DropKeys *[]string `json:"drop_keys,omitempty"`

	// 字段重命名源字段名称列表
	SourceKeys *[]string `json:"source_keys,omitempty"`

	// 字段重命名替换的字段名称列表
	DestKeys *[]string `json:"dest_keys,omitempty"`

	// json解析深度，默认为1层最大4层
	ExpandDepth *int32 `json:"expand_depth,omitempty"`

	// json解析字段名称连接符
	ExpandConnector *string `json:"expand_connector,omitempty"`

	// 自定义时间时间格式
	SourceFormat *string `json:"source_format,omitempty"`

	// 自定义时间字段value
	SourceValue *string `json:"source_value,omitempty"`

	// 是否开启自定义时间的开关
	SetTime *bool `json:"set_time,omitempty"`

	// 日志过滤白名单规则 key长度最大为256，value最大长度为128，key不可以重复
	Include map[string]string `json:"include,omitempty"`

	// 日志过滤白名单规则 key长度最大为256，value最大长度为128，key不可以重复
	Exclude map[string]string `json:"exclude,omitempty"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}

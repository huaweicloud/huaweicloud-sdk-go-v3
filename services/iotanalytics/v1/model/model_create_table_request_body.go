package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTableRequestBody struct {

	// 标签。只能包含数字、英文字母、中文字符、下划线、中划线、逗号以及斜杠。长度为0~128。
	Tags *string `json:"tags,omitempty"`

	// 新增表名称。
	TableName string `json:"table_name"`

	// 新增表别名。只能包含数字、英文字母、中文字符、下划线以及中划线。长度为0~32。
	TableAlias *string `json:"table_alias,omitempty"`

	// 新增表的描述信息。
	Description *string `json:"description,omitempty"`

	// 新增表的描述信息。
	Columns []Column `json:"columns"`

	// 新增表的数据类型，目前支持：Parquet、CSV格式。
	DataType string `json:"data_type"`

	// 数据来源。来源类型有：pipeline, default. 默认为default.
	DataSource *string `json:"data_source,omitempty"`

	// 仅当数据来源为pipeline时使用。数据的Data Store ID.
	DataStoreId *string `json:"data_store_id,omitempty"`

	// 表数据是否包含表头。只有CSV类型数据具有该属性。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// 用户自定义数据分隔符。只有CSV类型数据具有该属性。
	Delimiter *string `json:"delimiter,omitempty"`

	// 用户自定义引用字符，默认为双引号（即“\"”）。只有CSV类型数据具有该属性。
	QuoteChar *string `json:"quote_char,omitempty"`

	// 用户自定义转义字符，默认为反斜杠（即\\）。只有CSV类型数据具有该属性。
	EscapeChar *string `json:"escape_char,omitempty"`

	// 用户自定义日期类型，默认格式为“yyyy-MM-dd”。日期格式字符定义详见表3。只有CSV和JSON类型数据具有该属性。
	DateFormat *string `json:"date_format,omitempty"`

	// 用户自定义时间类型。默认格式为“yyyy-MM-dd HH:mm:ss”。时间戳格式字符定义详见表3。只有CSV和JSON类型数据具有该属性。
	TimestampFormat *string `json:"timestamp_format,omitempty"`
}

func (o CreateTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTableRequestBody", string(data)}, " ")
}

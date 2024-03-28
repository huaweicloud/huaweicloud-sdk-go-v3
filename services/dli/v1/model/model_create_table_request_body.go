package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableRequestBody 创建表的请求body体。
type CreateTableRequestBody struct {

	// 新增表名称。
	TableName string `json:"table_name"`

	// 数据存储的地方，分VIEW视图，OBS表和DLI表。
	DataLocation string `json:"data_location"`

	// 新增表的描述信息。
	Description *string `json:"description,omitempty"`

	// OBS表和DLI表必选参数。新增表的列。
	Columns []Column `json:"columns"`

	// OBS表必选参数。新增OBS表数据的类型，目前支持：Parquet、ORC、CSV、JSON、Carbon和Avro格式。
	DataType *string `json:"data_type,omitempty"`

	// OBS表必选参数。新增OBS表数据的存储路径，必须是OBS上的路径，以s3a开头。
	DataPath *string `json:"data_path,omitempty"`

	// OBS表非必选参数。OBS表数据是否包含表头。只有CSV类型数据具有该属性。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// OBS表非必选参数。用户自定义数据分隔符。只有CSV类型数据具有该属性。
	Delimiter *string `json:"delimiter,omitempty"`

	// OBS表非必选参数。用户自定义引用字符，默认为双引号（即“\\\"”）。只有CSV类型数据具有该属性。
	QuoteChar *string `json:"quote_char,omitempty"`

	// OBS表非必选参数。用户自定义转义字符，默认为反斜杠（即\"\\\\\"）。只有CSV类型数据具有该属性。
	EscapeChar *string `json:"escape_char,omitempty"`

	// OBS表非必选参数。用户自定义日期类型，默认格式为“yyyy-MM-dd”。只有CSV和JSON类型数据具有该属性。
	DateFormat *string `json:"date_format,omitempty"`

	// OBS表非必选参数。用户自定义时间类型。默认格式为“yyyy-MM-dd HH:mm:ss”。只有CSV和JSON类型数据具有该属性。
	TimestampFormat *string `json:"timestamp_format,omitempty"`

	// VIEW视图必选参数，创建视图时的数据选择语句。语句中涉及表需要使用“表=库名.表名”的格式
	SelectStatement *string `json:"select_statement,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTableRequestBody", string(data)}, " ")
}

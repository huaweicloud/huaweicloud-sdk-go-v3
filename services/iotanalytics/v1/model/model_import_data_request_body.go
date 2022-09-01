package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportDataRequestBody struct {

	// 导入的数据路径（当前仅支持导入OBS上的数据，且OBS路径须以s3a开头）。必须先把该OBS桶添加到离线数据源。
	DataPath string `json:"data_path" xml:"data_path"`

	// 导入的数据类型（当前支持csv、parquet、orc、json、avro数据类型）。
	DataType string `json:"data_type" xml:"data_type"`

	// 表ID。
	TableId string `json:"table_id" xml:"table_id"`

	// 导入数据中的第一行数据是否包括列名，即表头。默认为“false”，表示不包括列名。导入CSV类型数据时可指定。
	WithColumnHeader *string `json:"with_column_header,omitempty" xml:"with_column_header"`

	// 用户自定义数据分隔符，默认为逗号。导入CSV类型数据时可指定。
	Delimiter *string `json:"delimiter,omitempty" xml:"delimiter"`

	// 用户自定义引用字符，默认为双引号。导入CSV类型数据时可指定。
	QuoteChar *string `json:"quote_char,omitempty" xml:"quote_char"`

	// 用户自定义转义字符，默认为反斜杠。导入CSV类型数据时可指定。
	EscapeChar *string `json:"escape_char,omitempty" xml:"escape_char"`

	// 指定特定的日期格式，默认为“yyyy-MM-dd”。日期格式字符定义详见表3。导入CSV及JSON类型数据时可指定。
	DateFormat *string `json:"date_format,omitempty" xml:"date_format"`

	// 作业执行过程中的bad records存储目录。设置该配置项后，bad records不会导入到目标表。
	BadRecordsPath *string `json:"bad_records_path,omitempty" xml:"bad_records_path"`

	// 指定特定的时间格式，默认为“yyyy-MM-dd HH:mm:ss”。时间格式字符定义详见表3。导入CSV及JSON类型数据时可指定。
	TimestampFormat *string `json:"timestamp_format,omitempty" xml:"timestamp_format"`

	// 计算资源ID。
	ComputingResourceId string `json:"computing_resource_id" xml:"computing_resource_id"`
}

func (o ImportDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataRequestBody struct{}"
	}

	return strings.Join([]string{"ImportDataRequestBody", string(data)}, " ")
}

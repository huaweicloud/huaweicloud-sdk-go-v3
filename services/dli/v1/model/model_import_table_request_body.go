package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportTableRequestBody 导入数据
type ImportTableRequestBody struct {

	// 导入的数据路径（当前仅支持导入OBS上的数据，且OBS路径须以s3a开头）。
	DataPath string `json:"data_path"`

	// 导入的数据类型（当前支持csv、parquet、orc、json、avro数据类型）。
	DataType string `json:"data_type"`

	// 导入表所属的数据库名称。
	DatabaseName string `json:"database_name"`

	// 导入表的名称。
	TableName string `json:"table_name"`

	// 导入数据中的第一行数据是否包括列名，即表头。默认为“false”，表示不包括列名。导入CSV类型数据时可指定。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// 用户自定义数据分隔符，默认为逗号。导入CSV类型数据时可指定。
	Delimiter *string `json:"delimiter,omitempty"`

	// 用户自定义引用字符，默认为双引号。导入CSV类型数据时可指定。
	QuoteChar *string `json:"quote_char,omitempty"`

	// 用户自定义转义字符，默认为反斜杠。导入CSV类型数据时可指定。
	EscapeChar *string `json:"escape_char,omitempty"`

	// 指定特定的日期格式，默认为“yyyy-MM-dd”。日期格式字符定义详见表3。导入CSV及JSON类型数据时可指定。
	DateFormat *string `json:"date_format,omitempty"`

	// 作业执行过程中的bad records存储目录。设置该配置项后，bad records不会导入到目标表。
	BadRecordsPath *string `json:"bad_records_path,omitempty"`

	// 指定特定的时间格式，默认为“yyyy-MM-dd HH:mm:ss”。时间格式字符定义详见表3。导入CSV及JSON类型数据时可指定。
	TimestampFormat *string `json:"timestamp_format,omitempty"`

	// 指定执行该任务的队列。若不指定队列，将采用default队列执行操作。
	QueueName *string `json:"queue_name,omitempty"`

	// 是否覆盖之前的数据
	Overwrite *bool `json:"overwrite,omitempty"`

	// 表示需要导入到哪个分区
	PartitionSpec *string `json:"partition_spec,omitempty"`

	// 用于定义此配置项的参数
	Conf *[]string `json:"conf,omitempty"`
}

func (o ImportTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTableRequestBody struct{}"
	}

	return strings.Join([]string{"ImportTableRequestBody", string(data)}, " ")
}

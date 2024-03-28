package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobDetailResponse Response Object
type ShowSqlJobDetailResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 提交作业的用户。
	Owner *string `json:"owner,omitempty"`

	// 作业开始的时间。是单位为“毫秒”的时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业运行时长，单位毫秒。
	Duration *int64 `json:"duration,omitempty"`

	// 导出数据或保存查询结果时，指定的导出模式。
	ExportMode *string `json:"export_mode,omitempty"`

	// 记录其操作的表所在的数据库名称。类型为Import和Export作业才有“database_name”属性。
	DatabaseName *string `json:"database_name,omitempty"`

	// 记录其操作的表名称。类型为Import和Export作业才有“table_name”属性。
	TableName *string `json:"table_name,omitempty"`

	// 导入或导出的文件路径。
	DataPath *string `json:"data_path,omitempty"`

	// 导入或导出的数据类型（当前仅支持csv）。
	DataType *string `json:"data_type,omitempty"`

	// 导入作业时，导入的数据是否包括列名。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// 导入作业时，用户自定义数据分隔符。
	Delimiter *string `json:"delimiter,omitempty"`

	// 导入作业时，用户自定义引用字符。
	QuoteChar *string `json:"quote_char,omitempty"`

	// 导入作业时，用户自定义转义字符。
	EscapeChar *string `json:"escape_char,omitempty"`

	// 导入作业时，指定表的日期格式。
	DateFormat *string `json:"date_format,omitempty"`

	// 导入作业时，指定表的时间格式
	TimestampFormat *string `json:"timestamp_format,omitempty"`

	// 导出作业时，用户指定的压缩方式。
	Compress *string `json:"compress,omitempty"`

	// 作业标签
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobDetailResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlJobRunDetail struct {
	// 作业类型。

	SqlType *string `json:"sql_type,omitempty"`
	// 作业开始的时间。时间格式为ISO日期时间格式yyyy-MM-dd'T'HH:mm:ss'Z'

	StartTime *string `json:"start_time,omitempty"`
	// 作业运行时长，单位毫秒。

	Duration *int64 `json:"duration,omitempty"`
	// Insert作业执行过程中扫描的记录条数。

	InputRowCount *int64 `json:"input_row_count,omitempty"`
	// Insert作业执行过程中扫描到的错误记录数。

	BadRowCount *int64 `json:"bad_row_count,omitempty"`
	// 作业执行过程中扫描文件的大小。

	InputSize *int64 `json:"input_size,omitempty"`
	// 当前作业返回的结果总条数或insert作业插入的总条数。

	ResultCount *int64 `json:"result_count,omitempty"`
	// 记录其操作的表名称。类型为Import和Export作业才有“table_name”属性。

	TableName *string `json:"table_name,omitempty"`
	// Import类型的作业，记录其导入的数据是否包括列名。

	WithColumnHeader *bool `json:"with_column_header,omitempty"`
	// SQL查询的相关列信息的Json字符串。

	Detail *string `json:"detail,omitempty"`
	// 作业执行的SQL语句。

	Statement *string `json:"statement,omitempty"`
	// 系统提示信息。运行失败时，失败原因。

	Message *string `json:"message,omitempty"`
}

func (o SqlJobRunDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobRunDetail struct{}"
	}

	return strings.Join([]string{"SqlJobRunDetail", string(data)}, " ")
}

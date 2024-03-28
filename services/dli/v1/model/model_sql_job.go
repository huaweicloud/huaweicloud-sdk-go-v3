package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlJob 作业信息。
type SqlJob struct {

	// 作业ID。
	JobId string `json:"job_id"`

	// 作业类型。
	JobType string `json:"job_type"`

	// 作业提交的队列。
	QueueName string `json:"queue_name"`

	// 提交作业的用户。
	Owner string `json:"owner"`

	// 作业开始的时间。是单位为“毫秒”的时间戳。
	StartTime int64 `json:"start_time"`

	// 作业运行时长，单位毫秒。
	Duration *int64 `json:"duration,omitempty"`

	// 此作业的当前状态，包含提交（LAUNCHING）、运行中（RUNNING）、完成（FINISHED）、失败（FAILED）、取消（CANCELLED）。
	Status SqlJobStatus `json:"status"`

	// Insert作业执行过程中扫描的记录条数。
	InputRowCount *int64 `json:"input_row_count,omitempty"`

	// Insert作业执行过程中扫描到的错误记录数。
	BadRowCount *int64 `json:"bad_row_count,omitempty"`

	// 作业执行过程中扫描文件的大小。
	InputSize int64 `json:"input_size"`

	// 当前作业返回的结果总条数或insert作业插入的总条数。
	ResultCount int64 `json:"result_count"`

	// 记录其操作的表所在的数据库名称。类型为Import和Export作业才有“database_name”属性。
	DatabaseName *string `json:"database_name,omitempty"`

	// 记录其操作的表名称。类型为Import和Export作业才有“table_name”属性。
	TableName *string `json:"table_name,omitempty"`

	// Import类型的作业，记录其导入的数据是否包括列名。
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// SQL查询的相关列信息的Json字符串。
	Detail string `json:"detail"`

	// 作业执行的SQL语句。
	Statement string `json:"statement"`

	// 作业标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统提示信息。
	Message *string `json:"message,omitempty"`

	// 作业结束的时间。是单位为“毫秒”的时间戳。
	EndTime *int64 `json:"end_time,omitempty"`

	// 作业的CPU累计使用量
	CpuCost *string `json:"cpu_cost,omitempty"`

	// 作业的输出字节数
	OutputByte *string `json:"output_byte,omitempty"`
}

func (o SqlJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJob struct{}"
	}

	return strings.Join([]string{"SqlJob", string(data)}, " ")
}

type SqlJobStatus struct {
	value string
}

type SqlJobStatusEnum struct {
	LAUNCHING SqlJobStatus
	RUNNING   SqlJobStatus
	FINISHED  SqlJobStatus
	FAILED    SqlJobStatus
	CANCELLED SqlJobStatus
}

func GetSqlJobStatusEnum() SqlJobStatusEnum {
	return SqlJobStatusEnum{
		LAUNCHING: SqlJobStatus{
			value: "LAUNCHING",
		},
		RUNNING: SqlJobStatus{
			value: "RUNNING",
		},
		FINISHED: SqlJobStatus{
			value: "FINISHED",
		},
		FAILED: SqlJobStatus{
			value: "FAILED",
		},
		CANCELLED: SqlJobStatus{
			value: "CANCELLED",
		},
	}
}

func (c SqlJobStatus) Value() string {
	return c.value
}

func (c SqlJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlJobStatusResponse Response Object
type ShowSqlJobStatusResponse struct {

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty"`

	// 作业提交的队列。
	QueueName *string `json:"queue_name,omitempty"`

	// 提交作业的用户。
	Owner *string `json:"owner,omitempty"`

	// 作业开始的时间。是单位为“毫秒”的时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业运行时长，单位毫秒。
	Duration *int64 `json:"duration,omitempty"`

	// 此作业的当前状态，包含提交（LAUNCHING）、运行中（RUNNING）、完成（FINISHED）、失败（FAILED）、取消（CANCELLED）。
	Status *ShowSqlJobStatusResponseStatus `json:"status,omitempty"`

	// Insert作业执行过程中扫描的记录条数。
	InputRowCount *int64 `json:"input_row_count,omitempty"`

	// Insert作业执行过程中扫描到的错误记录数。
	BadRowCount *int64 `json:"bad_row_count,omitempty"`

	// 作业执行过程中扫描文件的大小。
	InputSize *int64 `json:"input_size,omitempty"`

	// 当前作业返回的结果总条数或insert作业插入的总条数。
	ResultCount *int32 `json:"result_count,omitempty"`

	// 记录其操作的表所在的数据库名称。类型为Import和Export作业才有“database_name”属性。
	DatabaseName *string `json:"database_name,omitempty"`

	// 记录其操作的表名称。类型为Import和Export作业才有“table_name”属性。
	TableName *string `json:"table_name,omitempty"`

	// SQL查询的相关列信息的Json字符串。
	Detail *string `json:"detail,omitempty"`

	// SQL配置参数信息Json字符串。
	UserConf *string `json:"user_conf,omitempty"`

	// 查询结果OBS路径
	ResultPath *string `json:"result_path,omitempty"`

	// 查询作业执行计划OBS路径
	ExecutionDetailsPath *string `json:"execution_details_path,omitempty"`

	// 查询结果格式
	ResultFormat *string `json:"result_format,omitempty"`

	// 作业执行的SQL语句。
	Statement *string `json:"statement,omitempty"`

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 作业执行方式
	JobMode *string `json:"job_mode,omitempty"`

	// 作业标签
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobStatusResponse", string(data)}, " ")
}

type ShowSqlJobStatusResponseStatus struct {
	value string
}

type ShowSqlJobStatusResponseStatusEnum struct {
	LAUNCHING ShowSqlJobStatusResponseStatus
	RUNNING   ShowSqlJobStatusResponseStatus
	FINISHED  ShowSqlJobStatusResponseStatus
	FAILED    ShowSqlJobStatusResponseStatus
	CANCELLED ShowSqlJobStatusResponseStatus
}

func GetShowSqlJobStatusResponseStatusEnum() ShowSqlJobStatusResponseStatusEnum {
	return ShowSqlJobStatusResponseStatusEnum{
		LAUNCHING: ShowSqlJobStatusResponseStatus{
			value: "LAUNCHING",
		},
		RUNNING: ShowSqlJobStatusResponseStatus{
			value: "RUNNING",
		},
		FINISHED: ShowSqlJobStatusResponseStatus{
			value: "FINISHED",
		},
		FAILED: ShowSqlJobStatusResponseStatus{
			value: "FAILED",
		},
		CANCELLED: ShowSqlJobStatusResponseStatus{
			value: "CANCELLED",
		},
	}
}

func (c ShowSqlJobStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowSqlJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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

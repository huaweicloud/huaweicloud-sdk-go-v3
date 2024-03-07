package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Job struct {

	// 作业类型： - NORMAL_JOB：表/文件迁移。 - BATCH_JOB：整库迁移。 - SCENARIO_JOB：场景迁移。
	JobType JobJobType `json:"job_type"`

	// 源端连接类型
	FromConnectorName string `json:"from-connector-name"`

	ToConfigValues *ConfigValues `json:"to-config-values"`

	// 目的端连接名称
	ToLinkName string `json:"to-link-name"`

	DriverConfigValues *ConfigValues `json:"driver-config-values"`

	FromConfigValues *ConfigValues `json:"from-config-values"`

	// 目的端连接类型
	ToConnectorName string `json:"to-connector-name"`

	// 作业名称，长度在1到240个字符之间
	Name string `json:"name"`

	// 源连接名称
	FromLinkName string `json:"from-link-name"`

	// 创建的用户。
	CreationUser *string `json:"creation-user,omitempty"`

	// 作业创建的时间，单位：毫秒。
	CreationDate *int64 `json:"creation-date,omitempty"`

	// 作业最后更新的时间，单位：毫秒。
	UpdateDate *int64 `json:"update-date,omitempty"`

	// 是否增量
	IsIncreJob *bool `json:"is_incre_job,omitempty"`

	// 标记
	Flag *int32 `json:"flag,omitempty"`

	// 已读文件数
	FilesRead *int32 `json:"files_read,omitempty"`

	// 作业最后更新的用户。
	UpdateUser *string `json:"update-user,omitempty"`

	// 外部ID。
	ExternalId *string `json:"external_id,omitempty"`

	// 作业类型
	Type *string `json:"type,omitempty"`

	// 执行_开始_日期。
	ExecuteStartDate *int64 `json:"execute_start_date,omitempty"`

	// 删除行数
	DeleteRows *int32 `json:"delete_rows,omitempty"`

	// 是否激活连接
	Enabled *bool `json:"enabled,omitempty"`

	// 写入字节
	BytesWritten *int64 `json:"bytes_written,omitempty"`

	// 作业ID
	Id *int32 `json:"id,omitempty"`

	// 用户是否使用sql
	IsUseSql *bool `json:"is_use_sql,omitempty"`

	// 更新行数
	UpdateRows *int32 `json:"update_rows,omitempty"`

	// 组_名称
	GroupName *string `json:"group_name,omitempty"`

	// 读取字节
	BytesRead *int64 `json:"bytes_read,omitempty"`

	// 执行_更新_日期。
	ExecuteUpdateDate *int64 `json:"execute_update_date,omitempty"`

	// 写入数据行数
	WriteRows *int32 `json:"write_rows,omitempty"`

	// 写入行数
	RowsWritten *int32 `json:"rows_written,omitempty"`

	// 读取的行数
	RowsRead *int64 `json:"rows_read,omitempty"`

	// 写入文件数
	FilesWritten *int32 `json:"files_written,omitempty"`

	// 是否增量
	IsIncrementing *bool `json:"is_incrementing,omitempty"`

	// 执行_创建_日期
	ExecuteCreateDate *int64 `json:"execute_create_date,omitempty"`

	// 作业最后的执行状态： - BOOTING：启动中。 - RUNNING：运行中。 - SUCCEEDED：成功。 - FAILED：失败。 - NEW：未被执行。
	Status *string `json:"status,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}

type JobJobType struct {
	value string
}

type JobJobTypeEnum struct {
	NORMAL_JOB   JobJobType
	BATCH_JOB    JobJobType
	SCENARIO_JOB JobJobType
}

func GetJobJobTypeEnum() JobJobTypeEnum {
	return JobJobTypeEnum{
		NORMAL_JOB: JobJobType{
			value: "NORMAL_JOB",
		},
		BATCH_JOB: JobJobType{
			value: "BATCH_JOB",
		},
		SCENARIO_JOB: JobJobType{
			value: "SCENARIO_JOB",
		},
	}
}

func (c JobJobType) Value() string {
	return c.value
}

func (c JobJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobJobType) UnmarshalJSON(b []byte) error {
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

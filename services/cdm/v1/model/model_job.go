package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Job struct {

	// 作业类型： - NORMAL_JOB：表/文件迁移。 - BATCH_JOB：整库迁移。 - SCENARIO_JOB：场景迁移。
	JobType *JobJobType `json:"job_type,omitempty" xml:"job_type"`

	// 源端连接类型
	FromConnectorName string `json:"from-connector-name" xml:"from-connector-name"`

	ToConfigValues *ConfigValues `json:"to-config-values" xml:"to-config-values"`

	// 目的端连接名称
	ToLinkName string `json:"to-link-name" xml:"to-link-name"`

	DriverConfigValues *ConfigValues `json:"driver-config-values" xml:"driver-config-values"`

	FromConfigValues *ConfigValues `json:"from-config-values" xml:"from-config-values"`

	// 目的端连接类型
	ToConnectorName *string `json:"to-connector-name,omitempty" xml:"to-connector-name"`

	// 作业名称，长度在1到240个字符之间
	Name *string `json:"name,omitempty" xml:"name"`

	// 源连接名称
	FromLinkName *string `json:"from-link-name,omitempty" xml:"from-link-name"`

	// 创建的用户。
	CreationUser *string `json:"creation-user,omitempty" xml:"creation-user"`

	// 作业创建的时间，单位：毫秒。
	CreationDate *int64 `json:"creation-date,omitempty" xml:"creation-date"`

	// 作业最后更新的时间，单位：毫秒。
	UpdateDate *int64 `json:"update-date,omitempty" xml:"update-date"`

	// 是否增量
	IsIncreJob *bool `json:"is_incre_job,omitempty" xml:"is_incre_job"`

	// 标记
	Flag *int32 `json:"flag,omitempty" xml:"flag"`

	// 已读文件数
	FilesRead *int32 `json:"files_read,omitempty" xml:"files_read"`

	// 作业最后更新的用户。
	UpdateUser *string `json:"update-user,omitempty" xml:"update-user"`

	// 外部ID。
	ExternalId *string `json:"external_id,omitempty" xml:"external_id"`

	// 作业类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 执行_开始_日期。
	ExecuteStartDate *int64 `json:"execute_start_date,omitempty" xml:"execute_start_date"`

	// 删除行数
	DeleteRows *int32 `json:"delete_rows,omitempty" xml:"delete_rows"`

	// 是否激活连接
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// 写入字节
	BytesWritten *int64 `json:"bytes_written,omitempty" xml:"bytes_written"`

	// 作业ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 用户是否使用sql
	IsUseSql *bool `json:"is_use_sql,omitempty" xml:"is_use_sql"`

	// 更新行数
	UpdateRows *int32 `json:"update_rows,omitempty" xml:"update_rows"`

	// 组_名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 读取字节
	BytesRead *int64 `json:"bytes_read,omitempty" xml:"bytes_read"`

	// 执行_更新_日期。
	ExecuteUpdateDate *int64 `json:"execute_update_date,omitempty" xml:"execute_update_date"`

	// 写入数据行数
	WriteRows *int32 `json:"write_rows,omitempty" xml:"write_rows"`

	// 写入行数
	RowsWritten *int32 `json:"rows_written,omitempty" xml:"rows_written"`

	// 读取的行数
	RowsRead *int64 `json:"rows_read,omitempty" xml:"rows_read"`

	// 写入文件数
	FilesWritten *int32 `json:"files_written,omitempty" xml:"files_written"`

	// 是否增量
	IsIncrementing *bool `json:"is_incrementing,omitempty" xml:"is_incrementing"`

	// 执行_创建_日期
	ExecuteCreateDate *int64 `json:"execute_create_date,omitempty" xml:"execute_create_date"`

	// 作业最后的执行状态： - BOOTING：启动中。 - RUNNING：运行中。 - SUCCEEDED：成功。 - FAILED：失败。 - NEW：未被执行。
	Status *string `json:"status,omitempty" xml:"status"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Job struct {
	// 作业类型： - NORMAL_JOB：表/文件迁移。 - BATCH_JOB：整库迁移。 - SCENARIO_JOB：场景迁移。

	JobType *JobJobType `json:"job_type,omitempty"`
	// 源端连接类型

	FromConnectorName *string `json:"from-connector-name,omitempty"`

	ToConfigValues *ConfigValues `json:"to-config-values,omitempty"`
	// 目的端连接名称

	ToLinkName *string `json:"to-link-name,omitempty"`

	DriverConfigValues *ConfigValues `json:"driver-config-values,omitempty"`

	FromConfigValues *ConfigValues `json:"from-config-values,omitempty"`
	// 目的端连接类型

	ToConnectorName *string `json:"to-connector-name,omitempty"`
	// 作业名称，长度在1到240个字符之间

	Name *string `json:"name,omitempty"`
	// 源连接名称

	FromLinkName *string `json:"from-link-name,omitempty"`
	// 创建的用户。

	CreationUser *string `json:"creation-user,omitempty"`
	// 作业创建的时间，单位：毫秒。

	CreationDate *int64 `json:"creation-date,omitempty"`
	// 作业最后更新的时间，单位：毫秒。

	UpdateDate *int64 `json:"update-date,omitempty"`
	// 作业最后更新的用户。

	UpdateUser *string `json:"update-user,omitempty"`
	// 作业最后的执行状态： - BOOTING：启动中。 - RUNNING：运行中。 - SUCCEEDED：成功。 - FAILED：失败。 - NEW：未被执行。

	Status *string `json:"status,omitempty"`
}

func (o Job) String() string {
	data, err := json.Marshal(o)
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

func (c JobJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
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

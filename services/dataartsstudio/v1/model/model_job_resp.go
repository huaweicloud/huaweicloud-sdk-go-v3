package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobResp 作业对象
type JobResp struct {

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 作业类型:  - REAL_TIME: 实时处理  - BATCH: 批处理
	JobType *JobRespJobType `json:"job_type,omitempty"`

	// 最近实例的运行结束时间
	LastInstanceEndTime *int64 `json:"last_instance_end_time,omitempty"`

	// 最近实例的运行状态
	LastInstanceStatus *string `json:"last_instance_status,omitempty"`

	// 最后更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 最后修改人
	LastUpdateUser *string `json:"last_update_user,omitempty"`

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 作业目录路径
	Path *string `json:"path,omitempty"`

	// 作业优先级
	Priority *int32 `json:"priority,omitempty"`

	// 单算子作业标识
	SingleNodeJobFlag *bool `json:"single_node_job_flag,omitempty"`

	// 作业启动时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`
}

func (o JobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResp struct{}"
	}

	return strings.Join([]string{"JobResp", string(data)}, " ")
}

type JobRespJobType struct {
	value string
}

type JobRespJobTypeEnum struct {
	REAL_TIME JobRespJobType
	BATCH     JobRespJobType
}

func GetJobRespJobTypeEnum() JobRespJobTypeEnum {
	return JobRespJobTypeEnum{
		REAL_TIME: JobRespJobType{
			value: "REAL_TIME",
		},
		BATCH: JobRespJobType{
			value: "BATCH",
		},
	}
}

func (c JobRespJobType) Value() string {
	return c.value
}

func (c JobRespJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobRespJobType) UnmarshalJSON(b []byte) error {
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

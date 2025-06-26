package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobDetail struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 相关实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 状态
	Status *JobDetailStatus `json:"status,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 实例创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 实例更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o JobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}

type JobDetailStatus struct {
	value string
}

type JobDetailStatusEnum struct {
	CREATING     JobDetailStatus
	INITIALIZING JobDetailStatus
	RUNNING      JobDetailStatus
	FAILED       JobDetailStatus
	SUCCESS      JobDetailStatus
}

func GetJobDetailStatusEnum() JobDetailStatusEnum {
	return JobDetailStatusEnum{
		CREATING: JobDetailStatus{
			value: "Creating",
		},
		INITIALIZING: JobDetailStatus{
			value: "Initializing",
		},
		RUNNING: JobDetailStatus{
			value: "Running",
		},
		FAILED: JobDetailStatus{
			value: "Failed",
		},
		SUCCESS: JobDetailStatus{
			value: "Success",
		},
	}
}

func (c JobDetailStatus) Value() string {
	return c.value
}

func (c JobDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobDetailStatus) UnmarshalJSON(b []byte) error {
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

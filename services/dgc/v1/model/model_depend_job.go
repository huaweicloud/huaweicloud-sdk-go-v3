package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DependJob struct {

	// 依赖的作业名称列表，必须依赖已存在的作业.
	Jobs []string `json:"jobs"`

	// 依赖周期
	DependPeriod *string `json:"dependPeriod,omitempty"`

	// 依赖作业任务执行失败处理策略
	DependFailPolicy *DependJobDependFailPolicy `json:"dependFailPolicy,omitempty"`

	// 依赖本工作空间的作业名称列表
	SameWorkSpaceJobs *[]DependWorkSpaceJob `json:"sameWorkSpaceJobs,omitempty"`

	// 依赖其他工作空间的作业名称列表
	OtherWorkSpaceJobs *[]DependWorkSpaceJob `json:"otherWorkSpaceJobs,omitempty"`
}

func (o DependJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependJob struct{}"
	}

	return strings.Join([]string{"DependJob", string(data)}, " ")
}

type DependJobDependFailPolicy struct {
	value string
}

type DependJobDependFailPolicyEnum struct {
	FAIL    DependJobDependFailPolicy
	IGNORE  DependJobDependFailPolicy
	SUSPEND DependJobDependFailPolicy
}

func GetDependJobDependFailPolicyEnum() DependJobDependFailPolicyEnum {
	return DependJobDependFailPolicyEnum{
		FAIL: DependJobDependFailPolicy{
			value: "FAIL",
		},
		IGNORE: DependJobDependFailPolicy{
			value: "IGNORE",
		},
		SUSPEND: DependJobDependFailPolicy{
			value: "SUSPEND",
		},
	}
}

func (c DependJobDependFailPolicy) Value() string {
	return c.value
}

func (c DependJobDependFailPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DependJobDependFailPolicy) UnmarshalJSON(b []byte) error {
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

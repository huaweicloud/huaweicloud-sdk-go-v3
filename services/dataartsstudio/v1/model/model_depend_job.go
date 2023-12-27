package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DependJob 依赖作业列表
type DependJob struct {

	// 依赖的作业名称列表，必须依赖已存在的作业。
	Jobs []string `json:"jobs"`

	// 依赖周期： - SAME_PERIOD：依赖被依赖作业的同周期任务的执行结果。 - PRE_PERIOD：依赖被依赖作业的前一周期任务的执行结果。
	DependPeriod *DependJobDependPeriod `json:"depend_period,omitempty"`

	// 依赖作业任务执行失败处理策略: - FAIL：停止作业，设置作业为失败状态。 - IGNORE：继续执行作业。 - SUSPEND： 挂起作业。
	DependFailPolicy *DependJobDependFailPolicy `json:"depend_fail_policy,omitempty"`
}

func (o DependJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependJob struct{}"
	}

	return strings.Join([]string{"DependJob", string(data)}, " ")
}

type DependJobDependPeriod struct {
	value string
}

type DependJobDependPeriodEnum struct {
	SAME_PERIOD DependJobDependPeriod
	PRE_PERIOD  DependJobDependPeriod
}

func GetDependJobDependPeriodEnum() DependJobDependPeriodEnum {
	return DependJobDependPeriodEnum{
		SAME_PERIOD: DependJobDependPeriod{
			value: "SAME_PERIOD",
		},
		PRE_PERIOD: DependJobDependPeriod{
			value: "PRE_PERIOD",
		},
	}
}

func (c DependJobDependPeriod) Value() string {
	return c.value
}

func (c DependJobDependPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DependJobDependPeriod) UnmarshalJSON(b []byte) error {
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

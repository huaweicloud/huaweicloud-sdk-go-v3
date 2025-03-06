package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateScheduleTaskRequestBody struct {

	// 实例ID列表。
	InstanceIds []string `json:"instance_ids"`

	// 任务开始时间。
	StartTime string `json:"start_time"`

	// 实例升级类型。
	UpgradeType CreateScheduleTaskRequestBodyUpgradeType `json:"upgrade_type"`

	// 实例升级操作。
	UpgradeAction CreateScheduleTaskRequestBodyUpgradeAction `json:"upgrade_action"`

	// 实例升级目标版本。
	TargetVersion string `json:"target_version"`
}

func (o CreateScheduleTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateScheduleTaskRequestBody", string(data)}, " ")
}

type CreateScheduleTaskRequestBodyUpgradeType struct {
	value string
}

type CreateScheduleTaskRequestBodyUpgradeTypeEnum struct {
	HOTFIX CreateScheduleTaskRequestBodyUpgradeType
}

func GetCreateScheduleTaskRequestBodyUpgradeTypeEnum() CreateScheduleTaskRequestBodyUpgradeTypeEnum {
	return CreateScheduleTaskRequestBodyUpgradeTypeEnum{
		HOTFIX: CreateScheduleTaskRequestBodyUpgradeType{
			value: "hotfix",
		},
	}
}

func (c CreateScheduleTaskRequestBodyUpgradeType) Value() string {
	return c.value
}

func (c CreateScheduleTaskRequestBodyUpgradeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScheduleTaskRequestBodyUpgradeType) UnmarshalJSON(b []byte) error {
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

type CreateScheduleTaskRequestBodyUpgradeAction struct {
	value string
}

type CreateScheduleTaskRequestBodyUpgradeActionEnum struct {
	UPGRADE_AUTO_COMMIT CreateScheduleTaskRequestBodyUpgradeAction
}

func GetCreateScheduleTaskRequestBodyUpgradeActionEnum() CreateScheduleTaskRequestBodyUpgradeActionEnum {
	return CreateScheduleTaskRequestBodyUpgradeActionEnum{
		UPGRADE_AUTO_COMMIT: CreateScheduleTaskRequestBodyUpgradeAction{
			value: "upgradeAutoCommit",
		},
	}
}

func (c CreateScheduleTaskRequestBodyUpgradeAction) Value() string {
	return c.value
}

func (c CreateScheduleTaskRequestBodyUpgradeAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScheduleTaskRequestBodyUpgradeAction) UnmarshalJSON(b []byte) error {
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

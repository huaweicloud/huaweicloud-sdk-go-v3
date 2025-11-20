package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckTaskRisk struct {

	// **参数解释：** 风险项名称 **取值范围：** 不涉及
	RiskName string `json:"riskName"`

	// **参数解释：** 风险等级 **取值范围：** - Warning: 中危，允许跳过 - Fatal: 高危，不允许跳过
	Level CheckTaskRiskLevel `json:"level"`

	// **参数解释：** 风险项检查状态 **取值范围：** - Init: 风险项检查状态，初始化 - Running: 风险项检查状态，检查中 - Failed: 风险项检查状态，检查完成有风险 - Success: 风险项检查状态，检查完成无风险
	Status CheckTaskRiskStatus `json:"status"`

	// **参数解释：** 风险检查结果说明 **取值范围：** 不涉及
	Message *string `json:"message,omitempty"`
}

func (o CheckTaskRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTaskRisk struct{}"
	}

	return strings.Join([]string{"CheckTaskRisk", string(data)}, " ")
}

type CheckTaskRiskLevel struct {
	value string
}

type CheckTaskRiskLevelEnum struct {
	WARNING CheckTaskRiskLevel
	FATAL   CheckTaskRiskLevel
}

func GetCheckTaskRiskLevelEnum() CheckTaskRiskLevelEnum {
	return CheckTaskRiskLevelEnum{
		WARNING: CheckTaskRiskLevel{
			value: "Warning",
		},
		FATAL: CheckTaskRiskLevel{
			value: "Fatal",
		},
	}
}

func (c CheckTaskRiskLevel) Value() string {
	return c.value
}

func (c CheckTaskRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckTaskRiskLevel) UnmarshalJSON(b []byte) error {
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

type CheckTaskRiskStatus struct {
	value string
}

type CheckTaskRiskStatusEnum struct {
	INIT    CheckTaskRiskStatus
	RUNNING CheckTaskRiskStatus
	FAILED  CheckTaskRiskStatus
	SUCCESS CheckTaskRiskStatus
}

func GetCheckTaskRiskStatusEnum() CheckTaskRiskStatusEnum {
	return CheckTaskRiskStatusEnum{
		INIT: CheckTaskRiskStatus{
			value: "Init",
		},
		RUNNING: CheckTaskRiskStatus{
			value: "Running",
		},
		FAILED: CheckTaskRiskStatus{
			value: "Failed",
		},
		SUCCESS: CheckTaskRiskStatus{
			value: "Success",
		},
	}
}

func (c CheckTaskRiskStatus) Value() string {
	return c.value
}

func (c CheckTaskRiskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckTaskRiskStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateHaConfigurationReqBody struct {

	// 高可用选项，接受高可用或者拒绝高可用功能。
	DisasterRecoveryChoice UpdateHaConfigurationReqBodyDisasterRecoveryChoice `json:"disaster_recovery_choice"`
}

func (o UpdateHaConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHaConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateHaConfigurationReqBody", string(data)}, " ")
}

type UpdateHaConfigurationReqBodyDisasterRecoveryChoice struct {
	value string
}

type UpdateHaConfigurationReqBodyDisasterRecoveryChoiceEnum struct {
	ACCEPT UpdateHaConfigurationReqBodyDisasterRecoveryChoice
	REJECT UpdateHaConfigurationReqBodyDisasterRecoveryChoice
}

func GetUpdateHaConfigurationReqBodyDisasterRecoveryChoiceEnum() UpdateHaConfigurationReqBodyDisasterRecoveryChoiceEnum {
	return UpdateHaConfigurationReqBodyDisasterRecoveryChoiceEnum{
		ACCEPT: UpdateHaConfigurationReqBodyDisasterRecoveryChoice{
			value: "ACCEPT",
		},
		REJECT: UpdateHaConfigurationReqBodyDisasterRecoveryChoice{
			value: "REJECT",
		},
	}
}

func (c UpdateHaConfigurationReqBodyDisasterRecoveryChoice) Value() string {
	return c.value
}

func (c UpdateHaConfigurationReqBodyDisasterRecoveryChoice) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHaConfigurationReqBodyDisasterRecoveryChoice) UnmarshalJSON(b []byte) error {
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

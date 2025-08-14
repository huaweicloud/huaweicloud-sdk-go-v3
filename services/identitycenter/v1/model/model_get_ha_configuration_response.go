package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetHaConfigurationResponse Response Object
type GetHaConfigurationResponse struct {

	// 高可用选项，接受高可用或者拒绝高可用功能。
	DisasterRecoveryChoice *GetHaConfigurationResponseDisasterRecoveryChoice `json:"disaster_recovery_choice,omitempty"`
	HttpStatusCode         int                                               `json:"-"`
}

func (o GetHaConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHaConfigurationResponse struct{}"
	}

	return strings.Join([]string{"GetHaConfigurationResponse", string(data)}, " ")
}

type GetHaConfigurationResponseDisasterRecoveryChoice struct {
	value string
}

type GetHaConfigurationResponseDisasterRecoveryChoiceEnum struct {
	ACCEPT GetHaConfigurationResponseDisasterRecoveryChoice
	REJECT GetHaConfigurationResponseDisasterRecoveryChoice
}

func GetGetHaConfigurationResponseDisasterRecoveryChoiceEnum() GetHaConfigurationResponseDisasterRecoveryChoiceEnum {
	return GetHaConfigurationResponseDisasterRecoveryChoiceEnum{
		ACCEPT: GetHaConfigurationResponseDisasterRecoveryChoice{
			value: "ACCEPT",
		},
		REJECT: GetHaConfigurationResponseDisasterRecoveryChoice{
			value: "REJECT",
		},
	}
}

func (c GetHaConfigurationResponseDisasterRecoveryChoice) Value() string {
	return c.value
}

func (c GetHaConfigurationResponseDisasterRecoveryChoice) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHaConfigurationResponseDisasterRecoveryChoice) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeTransactionSwitchStatusResponse Response Object
type ChangeTransactionSwitchStatusResponse struct {

	// 开关状态
	SwitchStatus   *ChangeTransactionSwitchStatusResponseSwitchStatus `json:"switch_status,omitempty"`
	HttpStatusCode int                                                `json:"-"`
}

func (o ChangeTransactionSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTransactionSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeTransactionSwitchStatusResponse", string(data)}, " ")
}

type ChangeTransactionSwitchStatusResponseSwitchStatus struct {
	value string
}

type ChangeTransactionSwitchStatusResponseSwitchStatusEnum struct {
	ENABLED  ChangeTransactionSwitchStatusResponseSwitchStatus
	DISABLED ChangeTransactionSwitchStatusResponseSwitchStatus
}

func GetChangeTransactionSwitchStatusResponseSwitchStatusEnum() ChangeTransactionSwitchStatusResponseSwitchStatusEnum {
	return ChangeTransactionSwitchStatusResponseSwitchStatusEnum{
		ENABLED: ChangeTransactionSwitchStatusResponseSwitchStatus{
			value: "Enabled",
		},
		DISABLED: ChangeTransactionSwitchStatusResponseSwitchStatus{
			value: "Disabled",
		},
	}
}

func (c ChangeTransactionSwitchStatusResponseSwitchStatus) Value() string {
	return c.value
}

func (c ChangeTransactionSwitchStatusResponseSwitchStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeTransactionSwitchStatusResponseSwitchStatus) UnmarshalJSON(b []byte) error {
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

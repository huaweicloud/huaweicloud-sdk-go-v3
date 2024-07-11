package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTransactionSwitchStatusResponse Response Object
type ShowTransactionSwitchStatusResponse struct {

	// 开关状态
	SwitchStatus   *ShowTransactionSwitchStatusResponseSwitchStatus `json:"switch_status,omitempty"`
	HttpStatusCode int                                              `json:"-"`
}

func (o ShowTransactionSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransactionSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowTransactionSwitchStatusResponse", string(data)}, " ")
}

type ShowTransactionSwitchStatusResponseSwitchStatus struct {
	value string
}

type ShowTransactionSwitchStatusResponseSwitchStatusEnum struct {
	ENABLED  ShowTransactionSwitchStatusResponseSwitchStatus
	DISABLED ShowTransactionSwitchStatusResponseSwitchStatus
}

func GetShowTransactionSwitchStatusResponseSwitchStatusEnum() ShowTransactionSwitchStatusResponseSwitchStatusEnum {
	return ShowTransactionSwitchStatusResponseSwitchStatusEnum{
		ENABLED: ShowTransactionSwitchStatusResponseSwitchStatus{
			value: "Enabled",
		},
		DISABLED: ShowTransactionSwitchStatusResponseSwitchStatus{
			value: "Disabled",
		},
	}
}

func (c ShowTransactionSwitchStatusResponseSwitchStatus) Value() string {
	return c.value
}

func (c ShowTransactionSwitchStatusResponseSwitchStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTransactionSwitchStatusResponseSwitchStatus) UnmarshalJSON(b []byte) error {
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

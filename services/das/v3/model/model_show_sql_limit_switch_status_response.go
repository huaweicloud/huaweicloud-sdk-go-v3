package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlLimitSwitchStatusResponse Response Object
type ShowSqlLimitSwitchStatusResponse struct {

	// 开关状态
	SwitchStatus   *ShowSqlLimitSwitchStatusResponseSwitchStatus `json:"switch_status,omitempty"`
	HttpStatusCode int                                           `json:"-"`
}

func (o ShowSqlLimitSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitSwitchStatusResponse", string(data)}, " ")
}

type ShowSqlLimitSwitchStatusResponseSwitchStatus struct {
	value string
}

type ShowSqlLimitSwitchStatusResponseSwitchStatusEnum struct {
	ON  ShowSqlLimitSwitchStatusResponseSwitchStatus
	OFF ShowSqlLimitSwitchStatusResponseSwitchStatus
}

func GetShowSqlLimitSwitchStatusResponseSwitchStatusEnum() ShowSqlLimitSwitchStatusResponseSwitchStatusEnum {
	return ShowSqlLimitSwitchStatusResponseSwitchStatusEnum{
		ON: ShowSqlLimitSwitchStatusResponseSwitchStatus{
			value: "ON",
		},
		OFF: ShowSqlLimitSwitchStatusResponseSwitchStatus{
			value: "OFF",
		},
	}
}

func (c ShowSqlLimitSwitchStatusResponseSwitchStatus) Value() string {
	return c.value
}

func (c ShowSqlLimitSwitchStatusResponseSwitchStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlLimitSwitchStatusResponseSwitchStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SetBalancerSwitchRequest struct {
	InstanceId string `json:"instance_id"`

	Action SetBalancerSwitchRequestAction `json:"action"`
}

func (o SetBalancerSwitchRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetBalancerSwitchRequest struct{}"
	}

	return strings.Join([]string{"SetBalancerSwitchRequest", string(data)}, " ")
}

type SetBalancerSwitchRequestAction struct {
	value string
}

type SetBalancerSwitchRequestActionEnum struct {
	START SetBalancerSwitchRequestAction
	STOP  SetBalancerSwitchRequestAction
}

func GetSetBalancerSwitchRequestActionEnum() SetBalancerSwitchRequestActionEnum {
	return SetBalancerSwitchRequestActionEnum{
		START: SetBalancerSwitchRequestAction{
			value: "start",
		},
		STOP: SetBalancerSwitchRequestAction{
			value: "stop",
		},
	}
}

func (c SetBalancerSwitchRequestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SetBalancerSwitchRequestAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

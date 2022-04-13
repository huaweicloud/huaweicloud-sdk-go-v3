package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ResetMessagesRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 重发类型。当前只支持“resend”。

	ActionId ResetMessagesRequestActionId `json:"action_id"`

	Body *ResetMessagesReq `json:"body,omitempty"`
}

func (o ResetMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessagesRequest struct{}"
	}

	return strings.Join([]string{"ResetMessagesRequest", string(data)}, " ")
}

type ResetMessagesRequestActionId struct {
	value string
}

type ResetMessagesRequestActionIdEnum struct {
	RESEND ResetMessagesRequestActionId
}

func GetResetMessagesRequestActionIdEnum() ResetMessagesRequestActionIdEnum {
	return ResetMessagesRequestActionIdEnum{
		RESEND: ResetMessagesRequestActionId{
			value: "resend",
		},
	}
}

func (c ResetMessagesRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetMessagesRequestActionId) UnmarshalJSON(b []byte) error {
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

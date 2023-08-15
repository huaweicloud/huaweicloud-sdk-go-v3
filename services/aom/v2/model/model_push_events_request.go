package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PushEventsRequest Request Object
type PushEventsRequest struct {

	// 告警所属的企业项目id。
	XEnterprisePrjectId *string `json:"x-enterprise-prject-id,omitempty"`

	// 接口请求动作。action=clear代表清除告警，不传或者传其他值默认为上报动作。
	Action *PushEventsRequestAction `json:"action,omitempty"`

	Body *EventList `json:"body,omitempty"`
}

func (o PushEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushEventsRequest struct{}"
	}

	return strings.Join([]string{"PushEventsRequest", string(data)}, " ")
}

type PushEventsRequestAction struct {
	value string
}

type PushEventsRequestActionEnum struct {
	CLEAR PushEventsRequestAction
}

func GetPushEventsRequestActionEnum() PushEventsRequestActionEnum {
	return PushEventsRequestActionEnum{
		CLEAR: PushEventsRequestAction{
			value: "clear",
		},
	}
}

func (c PushEventsRequestAction) Value() string {
	return c.value
}

func (c PushEventsRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PushEventsRequestAction) UnmarshalJSON(b []byte) error {
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

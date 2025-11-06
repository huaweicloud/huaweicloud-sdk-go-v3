package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PushEventsRequest Request Object
type PushEventsRequest struct {

	// 告警所属的企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 如果不传该参数值，默认为default企业项目，ID为0。
	EnterpriseProjectId *string `json:"enterprise-project-id,omitempty"`

	// 接口请求动作： - 不传或者传其他值：代表上报告警或事件动作。该参数值默认为空，即默认上报告警或事件。 - clear：代表清除告警动作。
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

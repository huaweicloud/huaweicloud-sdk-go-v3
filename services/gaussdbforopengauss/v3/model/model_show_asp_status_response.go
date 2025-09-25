package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAspStatusResponse Response Object
type ShowAspStatusResponse struct {

	// **参数解释**: 开关状态。 **取值范围**: - ON：开启ASP。 - OFF：关闭ASP。
	Status         *ShowAspStatusResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowAspStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAspStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAspStatusResponse", string(data)}, " ")
}

type ShowAspStatusResponseStatus struct {
	value string
}

type ShowAspStatusResponseStatusEnum struct {
	ON  ShowAspStatusResponseStatus
	OFF ShowAspStatusResponseStatus
}

func GetShowAspStatusResponseStatusEnum() ShowAspStatusResponseStatusEnum {
	return ShowAspStatusResponseStatusEnum{
		ON: ShowAspStatusResponseStatus{
			value: "ON",
		},
		OFF: ShowAspStatusResponseStatus{
			value: "OFF",
		},
	}
}

func (c ShowAspStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowAspStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAspStatusResponseStatus) UnmarshalJSON(b []byte) error {
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

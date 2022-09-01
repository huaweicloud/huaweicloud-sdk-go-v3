package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type StartGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id" xml:"graph_id"`

	// 图actionId
	ActionId StartGraphRequestActionId `json:"action_id" xml:"action_id"`

	Body *StartGraphReq `json:"body,omitempty" xml:"body"`
}

func (o StartGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartGraphRequest struct{}"
	}

	return strings.Join([]string{"StartGraphRequest", string(data)}, " ")
}

type StartGraphRequestActionId struct {
	value string
}

type StartGraphRequestActionIdEnum struct {
	START StartGraphRequestActionId
}

func GetStartGraphRequestActionIdEnum() StartGraphRequestActionIdEnum {
	return StartGraphRequestActionIdEnum{
		START: StartGraphRequestActionId{
			value: "start",
		},
	}
}

func (c StartGraphRequestActionId) Value() string {
	return c.value
}

func (c StartGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartGraphRequestActionId) UnmarshalJSON(b []byte) error {
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

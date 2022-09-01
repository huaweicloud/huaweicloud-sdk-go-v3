package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type RestartGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id" xml:"graph_id"`

	// 图actionId
	ActionId RestartGraphRequestActionId `json:"action_id" xml:"action_id"`
}

func (o RestartGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGraphRequest struct{}"
	}

	return strings.Join([]string{"RestartGraphRequest", string(data)}, " ")
}

type RestartGraphRequestActionId struct {
	value string
}

type RestartGraphRequestActionIdEnum struct {
	RESTART RestartGraphRequestActionId
}

func GetRestartGraphRequestActionIdEnum() RestartGraphRequestActionIdEnum {
	return RestartGraphRequestActionIdEnum{
		RESTART: RestartGraphRequestActionId{
			value: "restart",
		},
	}
}

func (c RestartGraphRequestActionId) Value() string {
	return c.value
}

func (c RestartGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestartGraphRequestActionId) UnmarshalJSON(b []byte) error {
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

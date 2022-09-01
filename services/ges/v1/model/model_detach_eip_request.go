package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DetachEipRequest struct {

	// 图ID。
	GraphId string `json:"graph_id" xml:"graph_id"`

	// 图actionId
	ActionId DetachEipRequestActionId `json:"action_id" xml:"action_id"`

	Body *UnbindEipReq `json:"body,omitempty" xml:"body"`
}

func (o DetachEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachEipRequest struct{}"
	}

	return strings.Join([]string{"DetachEipRequest", string(data)}, " ")
}

type DetachEipRequestActionId struct {
	value string
}

type DetachEipRequestActionIdEnum struct {
	UNBIND_EIP DetachEipRequestActionId
}

func GetDetachEipRequestActionIdEnum() DetachEipRequestActionIdEnum {
	return DetachEipRequestActionIdEnum{
		UNBIND_EIP: DetachEipRequestActionId{
			value: "unbindEip",
		},
	}
}

func (c DetachEipRequestActionId) Value() string {
	return c.value
}

func (c DetachEipRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetachEipRequestActionId) UnmarshalJSON(b []byte) error {
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

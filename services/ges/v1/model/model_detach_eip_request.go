package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DetachEipRequest Request Object
type DetachEipRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 图actionId
	ActionId DetachEipRequestActionId `json:"action_id"`

	Body *UnbindEipReq `json:"body,omitempty"`
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

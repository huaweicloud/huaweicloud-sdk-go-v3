package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type AttachEipRequest struct {
	// 图ID。

	GraphId string `json:"graph_id"`
	// 图actionId

	ActionId AttachEipRequestActionId `json:"action_id"`

	Body *BindEipReq `json:"body,omitempty"`
}

func (o AttachEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequest struct{}"
	}

	return strings.Join([]string{"AttachEipRequest", string(data)}, " ")
}

type AttachEipRequestActionId struct {
	value string
}

type AttachEipRequestActionIdEnum struct {
	BIND_EIP AttachEipRequestActionId
}

func GetAttachEipRequestActionIdEnum() AttachEipRequestActionIdEnum {
	return AttachEipRequestActionIdEnum{
		BIND_EIP: AttachEipRequestActionId{
			value: "bindEip",
		},
	}
}

func (c AttachEipRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachEipRequestActionId) UnmarshalJSON(b []byte) error {
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

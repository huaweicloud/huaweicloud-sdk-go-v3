package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmApprovalsRequest Request Object
type ConfirmApprovalsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 处理审批单结果类型
	ActionId ConfirmApprovalsRequestActionId `json:"action-id"`

	Body *ApprovalInfoParam `json:"body,omitempty"`
}

func (o ConfirmApprovalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApprovalsRequest struct{}"
	}

	return strings.Join([]string{"ConfirmApprovalsRequest", string(data)}, " ")
}

type ConfirmApprovalsRequestActionId struct {
	value string
}

type ConfirmApprovalsRequestActionIdEnum struct {
	REJECT  ConfirmApprovalsRequestActionId
	RESOLVE ConfirmApprovalsRequestActionId
}

func GetConfirmApprovalsRequestActionIdEnum() ConfirmApprovalsRequestActionIdEnum {
	return ConfirmApprovalsRequestActionIdEnum{
		REJECT: ConfirmApprovalsRequestActionId{
			value: "reject",
		},
		RESOLVE: ConfirmApprovalsRequestActionId{
			value: "resolve",
		},
	}
}

func (c ConfirmApprovalsRequestActionId) Value() string {
	return c.value
}

func (c ConfirmApprovalsRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmApprovalsRequestActionId) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApprovalStatusEnum 业务审批状态
type ApprovalStatusEnum struct {
	value string
}

type ApprovalStatusEnumEnum struct {
	DEVELOPING ApprovalStatusEnum
	APPROVED   ApprovalStatusEnum
	REJECT     ApprovalStatusEnum
	WITHDREW   ApprovalStatusEnum
}

func GetApprovalStatusEnumEnum() ApprovalStatusEnumEnum {
	return ApprovalStatusEnumEnum{
		DEVELOPING: ApprovalStatusEnum{
			value: "DEVELOPING",
		},
		APPROVED: ApprovalStatusEnum{
			value: "APPROVED",
		},
		REJECT: ApprovalStatusEnum{
			value: "REJECT",
		},
		WITHDREW: ApprovalStatusEnum{
			value: "WITHDREW",
		},
	}
}

func (c ApprovalStatusEnum) Value() string {
	return c.value
}

func (c ApprovalStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApprovalStatusEnum) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApprovedStateEnum 审批状态。 - APPROVING (审批中) - APPROVED (已审批) - UNAPPROVED (审批未通过)
type ApprovedStateEnum struct {
	value string
}

type ApprovedStateEnumEnum struct {
	APPROVING  ApprovedStateEnum
	APPROVED   ApprovedStateEnum
	UNAPPROVED ApprovedStateEnum
}

func GetApprovedStateEnumEnum() ApprovedStateEnumEnum {
	return ApprovedStateEnumEnum{
		APPROVING: ApprovedStateEnum{
			value: "APPROVING",
		},
		APPROVED: ApprovedStateEnum{
			value: "APPROVED",
		},
		UNAPPROVED: ApprovedStateEnum{
			value: "UNAPPROVED",
		},
	}
}

func (c ApprovedStateEnum) Value() string {
	return c.value
}

func (c ApprovedStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApprovedStateEnum) UnmarshalJSON(b []byte) error {
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

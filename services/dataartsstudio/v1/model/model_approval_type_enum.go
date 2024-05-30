package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApprovalTypeEnum 业务审核类型。 枚举值：   - PUBLISH: 发布   - OFFLINE: 下线
type ApprovalTypeEnum struct {
	value string
}

type ApprovalTypeEnumEnum struct {
	PUBLISH ApprovalTypeEnum
	OFFLINE ApprovalTypeEnum
}

func GetApprovalTypeEnumEnum() ApprovalTypeEnumEnum {
	return ApprovalTypeEnumEnum{
		PUBLISH: ApprovalTypeEnum{
			value: "PUBLISH",
		},
		OFFLINE: ApprovalTypeEnum{
			value: "OFFLINE",
		},
	}
}

func (c ApprovalTypeEnum) Value() string {
	return c.value
}

func (c ApprovalTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApprovalTypeEnum) UnmarshalJSON(b []byte) error {
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

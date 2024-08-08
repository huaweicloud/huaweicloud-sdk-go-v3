package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VerifyStatusEnum 审核状态： * `VERIFYING` - 审核中 * `VERIFY_PASS` - 审核通过 * `Verify_FAIL` - 审核不通过
type VerifyStatusEnum struct {
	value string
}

type VerifyStatusEnumEnum struct {
	VERIFYING   VerifyStatusEnum
	VERIFY_PASS VerifyStatusEnum
	VERIFY_FAIL VerifyStatusEnum
}

func GetVerifyStatusEnumEnum() VerifyStatusEnumEnum {
	return VerifyStatusEnumEnum{
		VERIFYING: VerifyStatusEnum{
			value: "VERIFYING",
		},
		VERIFY_PASS: VerifyStatusEnum{
			value: "VERIFY_PASS",
		},
		VERIFY_FAIL: VerifyStatusEnum{
			value: "Verify_FAIL",
		},
	}
}

func (c VerifyStatusEnum) Value() string {
	return c.value
}

func (c VerifyStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VerifyStatusEnum) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FindingType 访问分析结果类型。
type FindingType struct {
	value string
}

type FindingTypeEnum struct {
	EXTERNAL_ACCESS            FindingType
	UNUSED_IAM_USER_ACCESS_KEY FindingType
	UNUSED_IAM_USER_PASSWORD   FindingType
}

func GetFindingTypeEnum() FindingTypeEnum {
	return FindingTypeEnum{
		EXTERNAL_ACCESS: FindingType{
			value: "external_access",
		},
		UNUSED_IAM_USER_ACCESS_KEY: FindingType{
			value: "unused_iam_user_access_key",
		},
		UNUSED_IAM_USER_PASSWORD: FindingType{
			value: "unused_iam_user_password",
		},
	}
}

func (c FindingType) Value() string {
	return c.value
}

func (c FindingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FindingType) UnmarshalJSON(b []byte) error {
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

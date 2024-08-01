package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyDocumentType 校验策略类型。
type PolicyDocumentType struct {
	value string
}

type PolicyDocumentTypeEnum struct {
	IDENTITY_POLICY     PolicyDocumentType
	AGENCY_TRUST_POLICY PolicyDocumentType
	BUCKET_POLICY       PolicyDocumentType
}

func GetPolicyDocumentTypeEnum() PolicyDocumentTypeEnum {
	return PolicyDocumentTypeEnum{
		IDENTITY_POLICY: PolicyDocumentType{
			value: "identity_policy",
		},
		AGENCY_TRUST_POLICY: PolicyDocumentType{
			value: "agency_trust_policy",
		},
		BUCKET_POLICY: PolicyDocumentType{
			value: "bucket_policy",
		},
	}
}

func (c PolicyDocumentType) Value() string {
	return c.value
}

func (c PolicyDocumentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyDocumentType) UnmarshalJSON(b []byte) error {
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

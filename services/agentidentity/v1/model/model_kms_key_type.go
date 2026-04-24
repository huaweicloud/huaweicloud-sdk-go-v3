package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KmsKeyType The type of KMS key (CustomerManagedKey or ServiceManagedKey).
type KmsKeyType struct {
	value string
}

type KmsKeyTypeEnum struct {
	CUSTOMER_MANAGED_KEY KmsKeyType
	SERVICE_MANAGED_KEY  KmsKeyType
}

func GetKmsKeyTypeEnum() KmsKeyTypeEnum {
	return KmsKeyTypeEnum{
		CUSTOMER_MANAGED_KEY: KmsKeyType{
			value: "CustomerManagedKey",
		},
		SERVICE_MANAGED_KEY: KmsKeyType{
			value: "ServiceManagedKey",
		},
	}
}

func (c KmsKeyType) Value() string {
	return c.value
}

func (c KmsKeyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KmsKeyType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EncryptType 加密类型。 - KEEP：保持原有加密状态 - UNENCRYPTED：不加密 - ENCRYPTED：加密（需同步上传kms_id）
type EncryptType struct {
	value string
}

type EncryptTypeEnum struct {
	KEEP        EncryptType
	UNENCRYPTED EncryptType
	ENCRYPTED   EncryptType
}

func GetEncryptTypeEnum() EncryptTypeEnum {
	return EncryptTypeEnum{
		KEEP: EncryptType{
			value: "KEEP",
		},
		UNENCRYPTED: EncryptType{
			value: "UNENCRYPTED",
		},
		ENCRYPTED: EncryptType{
			value: "ENCRYPTED",
		},
	}
}

func (c EncryptType) Value() string {
	return c.value
}

func (c EncryptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptType) UnmarshalJSON(b []byte) error {
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

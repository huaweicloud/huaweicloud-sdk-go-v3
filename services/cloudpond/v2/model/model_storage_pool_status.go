package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StoragePoolStatus 存储池状态。 - CREATING：创建中 - AVAILABLE：可用 - EXPANDING：扩容中 - PENDING_PAYMENT：待支付 - FROZEN：已冻结
type StoragePoolStatus struct {
	value string
}

type StoragePoolStatusEnum struct {
	CREATING        StoragePoolStatus
	AVAILABLE       StoragePoolStatus
	EXPANDING       StoragePoolStatus
	PENDING_PAYMENT StoragePoolStatus
	FROZEN          StoragePoolStatus
}

func GetStoragePoolStatusEnum() StoragePoolStatusEnum {
	return StoragePoolStatusEnum{
		CREATING: StoragePoolStatus{
			value: "CREATING",
		},
		AVAILABLE: StoragePoolStatus{
			value: "AVAILABLE",
		},
		EXPANDING: StoragePoolStatus{
			value: "EXPANDING",
		},
		PENDING_PAYMENT: StoragePoolStatus{
			value: "PENDING_PAYMENT",
		},
		FROZEN: StoragePoolStatus{
			value: "FROZEN",
		},
	}
}

func (c StoragePoolStatus) Value() string {
	return c.value
}

func (c StoragePoolStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StoragePoolStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MaskStatus 屏蔽状态。UN_MASKED：未屏蔽，MASK_EFFECTIVE：已生效，MASK_INEFFECTIVE：未生效。
type MaskStatus struct {
	value string
}

type MaskStatusEnum struct {
	UN_MASKED        MaskStatus
	MASK_EFFECTIVE   MaskStatus
	MASK_INEFFECTIVE MaskStatus
}

func GetMaskStatusEnum() MaskStatusEnum {
	return MaskStatusEnum{
		UN_MASKED: MaskStatus{
			value: "UN_MASKED",
		},
		MASK_EFFECTIVE: MaskStatus{
			value: "MASK_EFFECTIVE",
		},
		MASK_INEFFECTIVE: MaskStatus{
			value: "MASK_INEFFECTIVE",
		},
	}
}

func (c MaskStatus) Value() string {
	return c.value
}

func (c MaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MaskStatus) UnmarshalJSON(b []byte) error {
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

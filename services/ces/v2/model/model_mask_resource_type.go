package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MaskResourceType 屏蔽资源类型。ALL_INSTANCE：全部资源，MULTI_INSTANCE：多实例资源
type MaskResourceType struct {
	value string
}

type MaskResourceTypeEnum struct {
	ALL_INSTANCE   MaskResourceType
	MULTI_INSTANCE MaskResourceType
}

func GetMaskResourceTypeEnum() MaskResourceTypeEnum {
	return MaskResourceTypeEnum{
		ALL_INSTANCE: MaskResourceType{
			value: "ALL_INSTANCE",
		},
		MULTI_INSTANCE: MaskResourceType{
			value: "MULTI_INSTANCE",
		},
	}
}

func (c MaskResourceType) Value() string {
	return c.value
}

func (c MaskResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MaskResourceType) UnmarshalJSON(b []byte) error {
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

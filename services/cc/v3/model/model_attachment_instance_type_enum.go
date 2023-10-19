package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AttachmentInstanceTypeEnum 接入网络实例类型，GDGW (专线)和ER_ROUTE_TABLE (路由表)。
type AttachmentInstanceTypeEnum struct {
	value string
}

type AttachmentInstanceTypeEnumEnum struct {
	GDGW           AttachmentInstanceTypeEnum
	ER_ROUTE_TABLE AttachmentInstanceTypeEnum
}

func GetAttachmentInstanceTypeEnumEnum() AttachmentInstanceTypeEnumEnum {
	return AttachmentInstanceTypeEnumEnum{
		GDGW: AttachmentInstanceTypeEnum{
			value: "GDGW",
		},
		ER_ROUTE_TABLE: AttachmentInstanceTypeEnum{
			value: "ER_ROUTE_TABLE",
		},
	}
}

func (c AttachmentInstanceTypeEnum) Value() string {
	return c.value
}

func (c AttachmentInstanceTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachmentInstanceTypeEnum) UnmarshalJSON(b []byte) error {
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

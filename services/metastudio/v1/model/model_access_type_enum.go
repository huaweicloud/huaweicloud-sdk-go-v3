package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessTypeEnum 平台对接类型。 * MEITUAN_OFFICIAL：对接美团直播平台官方接口 * MSS_STANDARD：对接MetaStudio定义的标准接口
type AccessTypeEnum struct {
	value string
}

type AccessTypeEnumEnum struct {
	MEITUAN_OFFICIAL AccessTypeEnum
	MSS_STANDARD     AccessTypeEnum
}

func GetAccessTypeEnumEnum() AccessTypeEnumEnum {
	return AccessTypeEnumEnum{
		MEITUAN_OFFICIAL: AccessTypeEnum{
			value: "MEITUAN_OFFICIAL",
		},
		MSS_STANDARD: AccessTypeEnum{
			value: "MSS_STANDARD",
		},
	}
}

func (c AccessTypeEnum) Value() string {
	return c.value
}

func (c AccessTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessTypeEnum) UnmarshalJSON(b []byte) error {
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

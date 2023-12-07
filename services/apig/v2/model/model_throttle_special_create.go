package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ThrottleSpecialCreate struct {

	// 流控时间内特殊对象能够访问API的最大次数限制
	CallLimits int32 `json:"call_limits"`

	// 特殊APP的编号或特殊租户的账号ID
	ObjectId string `json:"object_id"`

	// 特殊对象类型
	ObjectType ThrottleSpecialCreateObjectType `json:"object_type"`
}

func (o ThrottleSpecialCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleSpecialCreate struct{}"
	}

	return strings.Join([]string{"ThrottleSpecialCreate", string(data)}, " ")
}

type ThrottleSpecialCreateObjectType struct {
	value string
}

type ThrottleSpecialCreateObjectTypeEnum struct {
	APP  ThrottleSpecialCreateObjectType
	USER ThrottleSpecialCreateObjectType
}

func GetThrottleSpecialCreateObjectTypeEnum() ThrottleSpecialCreateObjectTypeEnum {
	return ThrottleSpecialCreateObjectTypeEnum{
		APP: ThrottleSpecialCreateObjectType{
			value: "APP",
		},
		USER: ThrottleSpecialCreateObjectType{
			value: "USER",
		},
	}
}

func (c ThrottleSpecialCreateObjectType) Value() string {
	return c.value
}

func (c ThrottleSpecialCreateObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThrottleSpecialCreateObjectType) UnmarshalJSON(b []byte) error {
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

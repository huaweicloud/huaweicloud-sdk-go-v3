package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListenerAccessControlType 访问控制策略类型,即黑名单还是白名单, 取值： - BLACK：黑名单 - WHITE：白名单
type ListenerAccessControlType struct {
	value string
}

type ListenerAccessControlTypeEnum struct {
	BLACK ListenerAccessControlType
	WHITE ListenerAccessControlType
}

func GetListenerAccessControlTypeEnum() ListenerAccessControlTypeEnum {
	return ListenerAccessControlTypeEnum{
		BLACK: ListenerAccessControlType{
			value: "BLACK",
		},
		WHITE: ListenerAccessControlType{
			value: "WHITE",
		},
	}
}

func (c ListenerAccessControlType) Value() string {
	return c.value
}

func (c ListenerAccessControlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerAccessControlType) UnmarshalJSON(b []byte) error {
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

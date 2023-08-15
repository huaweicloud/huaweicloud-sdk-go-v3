package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClientAffinity 客户端亲和性，取值： - SOURCE_IP：按源IP保持会话。 - NONE：关闭。
type ClientAffinity struct {
	value string
}

type ClientAffinityEnum struct {
	SOURCE_IP ClientAffinity
	NONE      ClientAffinity
}

func GetClientAffinityEnum() ClientAffinityEnum {
	return ClientAffinityEnum{
		SOURCE_IP: ClientAffinity{
			value: "SOURCE_IP",
		},
		NONE: ClientAffinity{
			value: "NONE",
		},
	}
}

func (c ClientAffinity) Value() string {
	return c.value
}

func (c ClientAffinity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClientAffinity) UnmarshalJSON(b []byte) error {
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

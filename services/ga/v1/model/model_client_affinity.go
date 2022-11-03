package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 客户端亲和性，取值： - SOURCE_IP：按源IP保持会话。 - NONE：关闭。
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

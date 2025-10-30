package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerType 服务器类型。 COMPUTE: 计算服务器 NETWORK: 网络服务器 BLOCK_STORAGE: 硬盘存储服务器
type ServerType struct {
	value string
}

type ServerTypeEnum struct {
	COMPUTE       ServerType
	NETWORK       ServerType
	BLOCK_STORAGE ServerType
}

func GetServerTypeEnum() ServerTypeEnum {
	return ServerTypeEnum{
		COMPUTE: ServerType{
			value: "COMPUTE",
		},
		NETWORK: ServerType{
			value: "NETWORK",
		},
		BLOCK_STORAGE: ServerType{
			value: "BLOCK_STORAGE",
		},
	}
}

func (c ServerType) Value() string {
	return c.value
}

func (c ServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerType) UnmarshalJSON(b []byte) error {
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

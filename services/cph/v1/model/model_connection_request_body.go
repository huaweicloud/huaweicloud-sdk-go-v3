package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConnectionRequestBody 获取云手机连接信息请求体
type ConnectionRequestBody struct {

	// phone_id数组，单次请求最大限制10
	PhoneIds []string `json:"phone_ids"`

	// 申请接入的客户端类型 - ANDROID: 安卓平台SDK - WINDOWS: Windows平台SDK - H5_MOBILE: H5移动端SDK - H5_PC: H5 PC端SDK - IOS: iOS平台SDK
	ClientType ConnectionRequestBodyClientType `json:"client_type"`
}

func (o ConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"ConnectionRequestBody", string(data)}, " ")
}

type ConnectionRequestBodyClientType struct {
	value string
}

type ConnectionRequestBodyClientTypeEnum struct {
	ANDROID   ConnectionRequestBodyClientType
	WINDOWS   ConnectionRequestBodyClientType
	H5_MOBILE ConnectionRequestBodyClientType
	H5_PC     ConnectionRequestBodyClientType
	IOS       ConnectionRequestBodyClientType
}

func GetConnectionRequestBodyClientTypeEnum() ConnectionRequestBodyClientTypeEnum {
	return ConnectionRequestBodyClientTypeEnum{
		ANDROID: ConnectionRequestBodyClientType{
			value: "ANDROID",
		},
		WINDOWS: ConnectionRequestBodyClientType{
			value: "WINDOWS",
		},
		H5_MOBILE: ConnectionRequestBodyClientType{
			value: "H5_MOBILE",
		},
		H5_PC: ConnectionRequestBodyClientType{
			value: "H5_PC",
		},
		IOS: ConnectionRequestBodyClientType{
			value: "IOS",
		},
	}
}

func (c ConnectionRequestBodyClientType) Value() string {
	return c.value
}

func (c ConnectionRequestBodyClientType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionRequestBodyClientType) UnmarshalJSON(b []byte) error {
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

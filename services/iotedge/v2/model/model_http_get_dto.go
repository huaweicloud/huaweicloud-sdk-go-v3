package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HttpGetDto struct {

	// 请求路径
	Path string `json:"path"`

	// 端口
	Port int32 `json:"port"`

	// 主机地址
	Host *string `json:"host,omitempty"`

	// 协议类型
	Scheme HttpGetDtoScheme `json:"scheme"`
}

func (o HttpGetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpGetDto struct{}"
	}

	return strings.Join([]string{"HttpGetDto", string(data)}, " ")
}

type HttpGetDtoScheme struct {
	value string
}

type HttpGetDtoSchemeEnum struct {
	HTTP  HttpGetDtoScheme
	HTTPS HttpGetDtoScheme
}

func GetHttpGetDtoSchemeEnum() HttpGetDtoSchemeEnum {
	return HttpGetDtoSchemeEnum{
		HTTP: HttpGetDtoScheme{
			value: "HTTP",
		},
		HTTPS: HttpGetDtoScheme{
			value: "HTTPS",
		},
	}
}

func (c HttpGetDtoScheme) Value() string {
	return c.value
}

func (c HttpGetDtoScheme) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HttpGetDtoScheme) UnmarshalJSON(b []byte) error {
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

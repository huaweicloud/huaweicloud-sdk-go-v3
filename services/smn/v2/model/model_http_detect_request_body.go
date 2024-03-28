package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HttpDetectRequestBody Http探测请求参数，包含protocol，endpoint，extension字段，详细限制见参数描述。
type HttpDetectRequestBody struct {

	// 协议类型，当前仅支持http/https，不可为空
	Protocol HttpDetectRequestBodyProtocol `json:"protocol"`

	// 待探测的终端地址，当前仅支持以http:// 或https://开头，不可为空
	Endpoint string `json:"endpoint"`

	Extension *HttpDetectRequestBodyExtension `json:"extension,omitempty"`
}

func (o HttpDetectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpDetectRequestBody struct{}"
	}

	return strings.Join([]string{"HttpDetectRequestBody", string(data)}, " ")
}

type HttpDetectRequestBodyProtocol struct {
	value string
}

type HttpDetectRequestBodyProtocolEnum struct {
	HTTP  HttpDetectRequestBodyProtocol
	HTTPS HttpDetectRequestBodyProtocol
}

func GetHttpDetectRequestBodyProtocolEnum() HttpDetectRequestBodyProtocolEnum {
	return HttpDetectRequestBodyProtocolEnum{
		HTTP: HttpDetectRequestBodyProtocol{
			value: "http",
		},
		HTTPS: HttpDetectRequestBodyProtocol{
			value: "https",
		},
	}
}

func (c HttpDetectRequestBodyProtocol) Value() string {
	return c.value
}

func (c HttpDetectRequestBodyProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HttpDetectRequestBodyProtocol) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExternalAccessProtocol 协议。
type ExternalAccessProtocol struct {
	value string
}

type ExternalAccessProtocolEnum struct {
	HTTP  ExternalAccessProtocol
	HTTPS ExternalAccessProtocol
}

func GetExternalAccessProtocolEnum() ExternalAccessProtocolEnum {
	return ExternalAccessProtocolEnum{
		HTTP: ExternalAccessProtocol{
			value: "HTTP",
		},
		HTTPS: ExternalAccessProtocol{
			value: "HTTPS",
		},
	}
}

func (c ExternalAccessProtocol) Value() string {
	return c.value
}

func (c ExternalAccessProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExternalAccessProtocol) UnmarshalJSON(b []byte) error {
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

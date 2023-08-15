package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProbeParameter struct {

	// type为http类型时生效
	Scheme *ProbeParameterScheme `json:"scheme,omitempty"`

	// type为http类型时生效。默认为POD的IP, 可以指定自定义的IP
	Host *string `json:"host,omitempty"`

	// type为http和tcp类型时生效。
	Port *int32 `json:"port,omitempty"`

	// type为http类型时生效。请求路径。
	Path *string `json:"path,omitempty"`

	// type为command类型时生效。命令列表
	Command *[]string `json:"command,omitempty"`
}

func (o ProbeParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProbeParameter struct{}"
	}

	return strings.Join([]string{"ProbeParameter", string(data)}, " ")
}

type ProbeParameterScheme struct {
	value string
}

type ProbeParameterSchemeEnum struct {
	HTTP  ProbeParameterScheme
	HTTPS ProbeParameterScheme
}

func GetProbeParameterSchemeEnum() ProbeParameterSchemeEnum {
	return ProbeParameterSchemeEnum{
		HTTP: ProbeParameterScheme{
			value: "HTTP",
		},
		HTTPS: ProbeParameterScheme{
			value: "HTTPS",
		},
	}
}

func (c ProbeParameterScheme) Value() string {
	return c.value
}

func (c ProbeParameterScheme) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProbeParameterScheme) UnmarshalJSON(b []byte) error {
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

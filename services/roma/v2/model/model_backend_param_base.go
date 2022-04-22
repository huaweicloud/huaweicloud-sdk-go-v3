package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackendParamBase struct {

	// 参数名称。 字符串由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。
	Name string `json:"name"`

	// 参数位置：PATH、QUERY、HEADER
	Location BackendParamBaseLocation `json:"location"`

	// 参数类别：REQUEST、CONSTANT、SYSTEM
	Origin BackendParamBaseOrigin `json:"origin"`

	// 参数值。字符长度不超过255 origin类别为REQUEST时，此字段值为req_params中的参数名称；  origin类别为CONSTANT时，此字段值为参数真正的值；  origin类别为SYSTEM时，此字段值为系统参数名称，系统参数分为网关内置参数、前端认证参数和后端认证参数，当api前端安全认证方式为自定义认证时，可以填写前端认证参数，当api开启后端认证时，可以填写后端认证参数。  网关内置参数取值及对应含义： - $context.sourceIp：API调用者的源地址 - $context.stage：API调用的部署环境 - $context.apiName: API的名称 - $context.apiId：API的ID - $context.appName: API调用者的APP对象名称 - $context.appId：API调用者的APP对象ID - $context.requestId：当次API调用生成跟踪ID - $context.serverAddr：网关的服务器地址 - $context.serverName：网关的服务器名称 - $context.handleTime：本次API调用的处理时间 - $context.providerAppId：API拥有者的应用对象ID，暂不支持使用  前端认证参数取值：以“$context.authorizer.frontend.”为前缀，如希望自定义认证校验通过返回的参数为aaa，那么此字段填写为$context.authorizer.frontend.aaa  后端认证参数取值：以“$context.authorizer.backend.”为前缀，如希望自定义认证校验通过返回的参数为aaa，那么此字段填写为$context.authorizer.backend.aaa
	Value string `json:"value"`

	// 描述。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`
}

func (o BackendParamBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendParamBase struct{}"
	}

	return strings.Join([]string{"BackendParamBase", string(data)}, " ")
}

type BackendParamBaseLocation struct {
	value string
}

type BackendParamBaseLocationEnum struct {
	PATH   BackendParamBaseLocation
	QUERY  BackendParamBaseLocation
	HEADER BackendParamBaseLocation
}

func GetBackendParamBaseLocationEnum() BackendParamBaseLocationEnum {
	return BackendParamBaseLocationEnum{
		PATH: BackendParamBaseLocation{
			value: "PATH",
		},
		QUERY: BackendParamBaseLocation{
			value: "QUERY",
		},
		HEADER: BackendParamBaseLocation{
			value: "HEADER",
		},
	}
}

func (c BackendParamBaseLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendParamBaseLocation) UnmarshalJSON(b []byte) error {
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

type BackendParamBaseOrigin struct {
	value string
}

type BackendParamBaseOriginEnum struct {
	REQUEST  BackendParamBaseOrigin
	CONSTANT BackendParamBaseOrigin
	SYSTEM   BackendParamBaseOrigin
}

func GetBackendParamBaseOriginEnum() BackendParamBaseOriginEnum {
	return BackendParamBaseOriginEnum{
		REQUEST: BackendParamBaseOrigin{
			value: "REQUEST",
		},
		CONSTANT: BackendParamBaseOrigin{
			value: "CONSTANT",
		},
		SYSTEM: BackendParamBaseOrigin{
			value: "SYSTEM",
		},
	}
}

func (c BackendParamBaseOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendParamBaseOrigin) UnmarshalJSON(b []byte) error {
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

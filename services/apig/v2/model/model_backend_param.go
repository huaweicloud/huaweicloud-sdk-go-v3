package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackendParam struct {

	// 参数类别：   - 后端服务参数：REQUEST   - 常量参数：CONSTANT   - 系统参数：SYSTEM
	Origin BackendParamOrigin `json:"origin"`

	// 参数名称。 字符串由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。
	Name string `json:"name"`

	// 描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 参数位置：PATH、QUERY、HEADER
	Location BackendParamLocation `json:"location"`

	// 参数值。字符长度不超过255 origin类别为REQUEST时，此字段值为req_params中的参数名称；  origin类别为CONSTANT时，此字段值为参数真正的值；  origin类别为SYSTEM时，此字段值为系统参数名称，系统参数分为网关内置参数、前端认证参数和后端认证参数，当api前端安全认证方式为自定义认证时，可以填写前端认证参数，当api开启后端认证时，可以填写后端认证参数。  网关内置参数取值及对应含义： - $context.sourceIp：API调用者的源地址 - $context.stage：API调用的部署环境 - $context.apiId：API的ID - $context.appId：API调用者的APP对象ID - $context.requestId：当次API调用生成请求ID - $context.serverAddr：网关的服务器地址 - $context.serverName：网关的服务器名称 - $context.handleTime：本次API调用的处理时间 - $context.providerAppId：API拥有者的应用对象ID，暂不支持使用 - $context.member_group_name: 后端服务器所属的服务器分组名称 - $context.member_group_id: 后端服务器所属的服务器分组ID  前端认证参数取值：以“$context.authorizer.frontend.”为前缀，如希望自定义认证校验通过返回的参数为aaa，那么此字段填写为$context.authorizer.frontend.aaa  后端认证参数取值：以“$context.authorizer.backend.”为前缀，如希望自定义认证校验通过返回的参数为aaa，那么此字段填写为$context.authorizer.backend.aaa
	Value string `json:"value"`

	// 参数编号
	Id *string `json:"id,omitempty"`

	// 对应的请求参数编号
	ReqParamId *string `json:"req_param_id,omitempty"`
}

func (o BackendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendParam struct{}"
	}

	return strings.Join([]string{"BackendParam", string(data)}, " ")
}

type BackendParamOrigin struct {
	value string
}

type BackendParamOriginEnum struct {
	REQUEST  BackendParamOrigin
	CONSTANT BackendParamOrigin
	SYSTEM   BackendParamOrigin
}

func GetBackendParamOriginEnum() BackendParamOriginEnum {
	return BackendParamOriginEnum{
		REQUEST: BackendParamOrigin{
			value: "REQUEST",
		},
		CONSTANT: BackendParamOrigin{
			value: "CONSTANT",
		},
		SYSTEM: BackendParamOrigin{
			value: "SYSTEM",
		},
	}
}

func (c BackendParamOrigin) Value() string {
	return c.value
}

func (c BackendParamOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendParamOrigin) UnmarshalJSON(b []byte) error {
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

type BackendParamLocation struct {
	value string
}

type BackendParamLocationEnum struct {
	PATH   BackendParamLocation
	QUERY  BackendParamLocation
	HEADER BackendParamLocation
}

func GetBackendParamLocationEnum() BackendParamLocationEnum {
	return BackendParamLocationEnum{
		PATH: BackendParamLocation{
			value: "PATH",
		},
		QUERY: BackendParamLocation{
			value: "QUERY",
		},
		HEADER: BackendParamLocation{
			value: "HEADER",
		},
	}
}

func (c BackendParamLocation) Value() string {
	return c.value
}

func (c BackendParamLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendParamLocation) UnmarshalJSON(b []byte) error {
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

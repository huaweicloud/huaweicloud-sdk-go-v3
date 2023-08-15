package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRedirectUrlConfig 转发到的url配置。  当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。  当action为REDIRECT_TO_URL时生效，且为必选字段，其他action不可指定，否则报错。  格式：protocol://host:port/path?query  protocol、host、port、path不允许同时不传或同时传${xxx} （${xxx}表示原值，如${host}表示被转发的请求URL的host部分）。 protocol和port传入的值不能与l7policy关联的监听器一致且host、path同时不传或同时传${xxx}。   [共享型负载均衡器下的转发策略不支持该字段，传入会报错。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt)
type CreateRedirectUrlConfig struct {

	// 重定向的协议。默认值${protocol}表示继承原值（即与被转发请求保持一致）。  取值范围： - HTTP - HTTPS - ${protocol}
	Protocol *CreateRedirectUrlConfigProtocol `json:"protocol,omitempty"`

	// 重定向的主机名。字符串只能包含英文字母、数字、“-”、“.”，必须以字母、数字开头。 默认值${host}表示继承原值（即与被转发请求保持一致）。
	Host *string `json:"host,omitempty"`

	// 重定向到的端口。默认值${port}表示继承原值（即与被转发请求保持一致）。
	Port *string `json:"port,omitempty"`

	// 重定向的路径。默认值${path}表示继承原值（即与被转发请求保持一致）。  只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:|\\/()\\[\\]{}，且必须以\"/\"开头。其中$1，$2会匹配请求url通配符星号（*）
	Path *string `json:"path,omitempty"`

	// 重定向的查询字符串。默认${query}表示继承原值（即与被转发请求保持一致）。举例如下：  若该字段被设置为：${query}&name=my_name，则在转发符合条件的URL （如https://www.xxx.com:8080/elb?type=loadbalancer， 此时${query}表示type=loadbalancer）时，将会重定向到 https://www.xxx.com:8080/elb?type=loadbalancer&name=my_name。  只能包含英文字母、数字和特殊字符：!$&'()*+,-./:;=?@^_`。字母区分大小写。其中$1，$2会匹配请求url通配符星号（*）
	Query *string `json:"query,omitempty"`

	// 重定向后的返回码。  取值范围： - 301 - 302 - 303 - 307 - 308
	StatusCode CreateRedirectUrlConfigStatusCode `json:"status_code"`
}

func (o CreateRedirectUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedirectUrlConfig struct{}"
	}

	return strings.Join([]string{"CreateRedirectUrlConfig", string(data)}, " ")
}

type CreateRedirectUrlConfigProtocol struct {
	value string
}

type CreateRedirectUrlConfigProtocolEnum struct {
	HTTP      CreateRedirectUrlConfigProtocol
	HTTPS     CreateRedirectUrlConfigProtocol
	_PROTOCOL CreateRedirectUrlConfigProtocol
}

func GetCreateRedirectUrlConfigProtocolEnum() CreateRedirectUrlConfigProtocolEnum {
	return CreateRedirectUrlConfigProtocolEnum{
		HTTP: CreateRedirectUrlConfigProtocol{
			value: "HTTP",
		},
		HTTPS: CreateRedirectUrlConfigProtocol{
			value: "HTTPS",
		},
		_PROTOCOL: CreateRedirectUrlConfigProtocol{
			value: "${protocol}",
		},
	}
}

func (c CreateRedirectUrlConfigProtocol) Value() string {
	return c.value
}

func (c CreateRedirectUrlConfigProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRedirectUrlConfigProtocol) UnmarshalJSON(b []byte) error {
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

type CreateRedirectUrlConfigStatusCode struct {
	value string
}

type CreateRedirectUrlConfigStatusCodeEnum struct {
	E_301 CreateRedirectUrlConfigStatusCode
	E_302 CreateRedirectUrlConfigStatusCode
	E_303 CreateRedirectUrlConfigStatusCode
	E_307 CreateRedirectUrlConfigStatusCode
	E_308 CreateRedirectUrlConfigStatusCode
}

func GetCreateRedirectUrlConfigStatusCodeEnum() CreateRedirectUrlConfigStatusCodeEnum {
	return CreateRedirectUrlConfigStatusCodeEnum{
		E_301: CreateRedirectUrlConfigStatusCode{
			value: "301",
		},
		E_302: CreateRedirectUrlConfigStatusCode{
			value: "302",
		},
		E_303: CreateRedirectUrlConfigStatusCode{
			value: "303",
		},
		E_307: CreateRedirectUrlConfigStatusCode{
			value: "307",
		},
		E_308: CreateRedirectUrlConfigStatusCode{
			value: "308",
		},
	}
}

func (c CreateRedirectUrlConfigStatusCode) Value() string {
	return c.value
}

func (c CreateRedirectUrlConfigStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRedirectUrlConfigStatusCode) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDomainConfigRequestBody struct {

	// 源站类型。 0 - 源站IP， 1 - 源站域名。
	RealServerType UpdateDomainConfigRequestBodyRealServerType `json:"real_server_type"`

	// 源站ip/源站域名
	RealServer string `json:"real_server"`

	// HTTP端口，与port_https不能同时为空。DDoS高防支持的HTTP端口可用ListDomainPort接口查询。
	PortHttp *[]int32 `json:"port_http,omitempty"`

	// HTTPS端口，与port_http不能同时为空。DDoS高防支持的HTTPS端口可用ListDomainPort接口查询。
	PortHttps *[]int32 `json:"port_https,omitempty"`

	// 防护区域：0-中国大陆、1-中国大陆外
	OverseasType *string `json:"overseas_type,omitempty"`

	// 证书名称
	CertName *string `json:"cert_name,omitempty"`
}

func (o UpdateDomainConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainConfigRequestBody", string(data)}, " ")
}

type UpdateDomainConfigRequestBodyRealServerType struct {
	value int32
}

type UpdateDomainConfigRequestBodyRealServerTypeEnum struct {
	E_1 UpdateDomainConfigRequestBodyRealServerType
	E_0 UpdateDomainConfigRequestBodyRealServerType
}

func GetUpdateDomainConfigRequestBodyRealServerTypeEnum() UpdateDomainConfigRequestBodyRealServerTypeEnum {
	return UpdateDomainConfigRequestBodyRealServerTypeEnum{
		E_1: UpdateDomainConfigRequestBodyRealServerType{
			value: 1,
		}, E_0: UpdateDomainConfigRequestBodyRealServerType{
			value: 0,
		},
	}
}

func (c UpdateDomainConfigRequestBodyRealServerType) Value() int32 {
	return c.value
}

func (c UpdateDomainConfigRequestBodyRealServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainConfigRequestBodyRealServerType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

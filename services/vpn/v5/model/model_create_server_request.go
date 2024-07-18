package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateServerRequest struct {

	// 隧道协议类型
	TunnelProtocol CreateServerRequestTunnelProtocol `json:"tunnel_protocol"`

	// 客户端网段
	ClientCidr string `json:"client_cidr"`

	// 本端网段列表，至少有一个本端网段
	LocalSubnets []string `json:"local_subnets"`

	// 客户端认证类型
	ClientAuthType CreateServerRequestClientAuthType `json:"client_auth_type"`

	ServerCertificate *CreateServerRequestServerCertificate `json:"server_certificate,omitempty"`

	// 客户端证书列表。隧道协议类型是SSL且认证方式是证书认证时，必填
	ClientCaCertificates *[]CreateServerRequestClientCaCertificates `json:"client_ca_certificates,omitempty"`

	SslOptions *CreateServerRequestSslOptions `json:"ssl_options,omitempty"`
}

func (o CreateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerRequest struct{}"
	}

	return strings.Join([]string{"CreateServerRequest", string(data)}, " ")
}

type CreateServerRequestTunnelProtocol struct {
	value string
}

type CreateServerRequestTunnelProtocolEnum struct {
	SSL CreateServerRequestTunnelProtocol
}

func GetCreateServerRequestTunnelProtocolEnum() CreateServerRequestTunnelProtocolEnum {
	return CreateServerRequestTunnelProtocolEnum{
		SSL: CreateServerRequestTunnelProtocol{
			value: "SSL",
		},
	}
}

func (c CreateServerRequestTunnelProtocol) Value() string {
	return c.value
}

func (c CreateServerRequestTunnelProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerRequestTunnelProtocol) UnmarshalJSON(b []byte) error {
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

type CreateServerRequestClientAuthType struct {
	value string
}

type CreateServerRequestClientAuthTypeEnum struct {
	CERT           CreateServerRequestClientAuthType
	LOCAL_PASSWORD CreateServerRequestClientAuthType
}

func GetCreateServerRequestClientAuthTypeEnum() CreateServerRequestClientAuthTypeEnum {
	return CreateServerRequestClientAuthTypeEnum{
		CERT: CreateServerRequestClientAuthType{
			value: "CERT",
		},
		LOCAL_PASSWORD: CreateServerRequestClientAuthType{
			value: "LOCAL_PASSWORD",
		},
	}
}

func (c CreateServerRequestClientAuthType) Value() string {
	return c.value
}

func (c CreateServerRequestClientAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerRequestClientAuthType) UnmarshalJSON(b []byte) error {
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

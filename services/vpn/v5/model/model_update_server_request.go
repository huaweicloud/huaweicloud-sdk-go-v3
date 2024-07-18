package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateServerRequest struct {

	// 客户端网段
	ClientCidr *string `json:"client_cidr,omitempty"`

	// 本端网段列表,至少有一个本端网段
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	ServerCertificate *UpdateServerRequestServerCertificate `json:"server_certificate,omitempty"`

	SslOptions *UpdateServerRequestSslOptions `json:"ssl_options,omitempty"`

	// 客户端认证类型
	ClientAuthType *UpdateServerRequestClientAuthType `json:"client_auth_type,omitempty"`
}

func (o UpdateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerRequest", string(data)}, " ")
}

type UpdateServerRequestClientAuthType struct {
	value string
}

type UpdateServerRequestClientAuthTypeEnum struct {
	CERT           UpdateServerRequestClientAuthType
	LOCAL_PASSWORD UpdateServerRequestClientAuthType
}

func GetUpdateServerRequestClientAuthTypeEnum() UpdateServerRequestClientAuthTypeEnum {
	return UpdateServerRequestClientAuthTypeEnum{
		CERT: UpdateServerRequestClientAuthType{
			value: "CERT",
		},
		LOCAL_PASSWORD: UpdateServerRequestClientAuthType{
			value: "LOCAL_PASSWORD",
		},
	}
}

func (c UpdateServerRequestClientAuthType) Value() string {
	return c.value
}

func (c UpdateServerRequestClientAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServerRequestClientAuthType) UnmarshalJSON(b []byte) error {
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

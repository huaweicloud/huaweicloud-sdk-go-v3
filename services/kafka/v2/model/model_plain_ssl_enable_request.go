package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PlainSslEnableRequest 修改Kafka实例的接入方式。
type PlainSslEnableRequest struct {

	// 需要开启或者关闭的接入方式。
	Protocol *PlainSslEnableRequestProtocol `json:"protocol,omitempty"`

	// - true：开启指定的接入方式。 - false：关闭指定的接入方式。
	Enable *bool `json:"enable,omitempty"`

	// 首次开启SASL时，需要输入用户名。 实例创建后，关闭SASL并不会删除已经创建的用户，再次开启SASL时无需传入用户名，传入的用户名将无效。
	UserName *string `json:"user_name,omitempty"`

	// 首次开启SASL时，需要输入用户名的密码。
	PassWord *string `json:"pass_word,omitempty"`

	// 开启SASL后使用的认证机制。仅在第一次开启SASL时传入生效。生效后再次传入无效。 - PLAIN：简单的用户名密码校验。 - SCRAM-SHA-512：用户凭证校验，安全性比PLAIN机制更高。
	SaslEnabledMechanisms *[]PlainSslEnableRequestSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`
}

func (o PlainSslEnableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlainSslEnableRequest struct{}"
	}

	return strings.Join([]string{"PlainSslEnableRequest", string(data)}, " ")
}

type PlainSslEnableRequestProtocol struct {
	value string
}

type PlainSslEnableRequestProtocolEnum struct {
	PRIVATE_PLAIN_ENABLE          PlainSslEnableRequestProtocol
	PRIVATE_SASL_SSL_ENABLE       PlainSslEnableRequestProtocol
	PRIVATE_SASL_PLAINTEXT_ENABLE PlainSslEnableRequestProtocol
	PUBLIC_PLAIN_ENABLE           PlainSslEnableRequestProtocol
	PUBLIC_SASL_SSL_ENABLE        PlainSslEnableRequestProtocol
	PUBLIC_SASL_PLAINTEXT_ENABLE  PlainSslEnableRequestProtocol
}

func GetPlainSslEnableRequestProtocolEnum() PlainSslEnableRequestProtocolEnum {
	return PlainSslEnableRequestProtocolEnum{
		PRIVATE_PLAIN_ENABLE: PlainSslEnableRequestProtocol{
			value: "private_plain_enable",
		},
		PRIVATE_SASL_SSL_ENABLE: PlainSslEnableRequestProtocol{
			value: "private_sasl_ssl_enable",
		},
		PRIVATE_SASL_PLAINTEXT_ENABLE: PlainSslEnableRequestProtocol{
			value: "private_sasl_plaintext_enable",
		},
		PUBLIC_PLAIN_ENABLE: PlainSslEnableRequestProtocol{
			value: "public_plain_enable",
		},
		PUBLIC_SASL_SSL_ENABLE: PlainSslEnableRequestProtocol{
			value: "public_sasl_ssl_enable",
		},
		PUBLIC_SASL_PLAINTEXT_ENABLE: PlainSslEnableRequestProtocol{
			value: "public_sasl_plaintext_enable",
		},
	}
}

func (c PlainSslEnableRequestProtocol) Value() string {
	return c.value
}

func (c PlainSslEnableRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PlainSslEnableRequestProtocol) UnmarshalJSON(b []byte) error {
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

type PlainSslEnableRequestSaslEnabledMechanisms struct {
	value string
}

type PlainSslEnableRequestSaslEnabledMechanismsEnum struct {
	SCRAM_SHA_512 PlainSslEnableRequestSaslEnabledMechanisms
	PLAIN         PlainSslEnableRequestSaslEnabledMechanisms
}

func GetPlainSslEnableRequestSaslEnabledMechanismsEnum() PlainSslEnableRequestSaslEnabledMechanismsEnum {
	return PlainSslEnableRequestSaslEnabledMechanismsEnum{
		SCRAM_SHA_512: PlainSslEnableRequestSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
		PLAIN: PlainSslEnableRequestSaslEnabledMechanisms{
			value: "PLAIN",
		},
	}
}

func (c PlainSslEnableRequestSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c PlainSslEnableRequestSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PlainSslEnableRequestSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

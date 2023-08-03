package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SourceKafkaMqParameters struct {

	// kafka消费组
	Group string `json:"group"`

	// kafka实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// kafka实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// kafka topic名称
	Topic string `json:"topic"`

	// 消费点位
	SeekTo *SourceKafkaMqParametersSeekTo `json:"seek_to,omitempty"`

	// SASL_SSL是否开启
	EnableSaslSsl *bool `json:"enable_sasl_ssl,omitempty"`

	// SASL认证机制
	SaslMechanism *SourceKafkaMqParametersSaslMechanism `json:"sasl_mechanism,omitempty"`

	// SASL证书地址，配置的obs地址
	SslCertificateUrl *string `json:"ssl_certificate_url,omitempty"`

	// SASL证书密码
	SslCertificatePwd *string `json:"ssl_certificate_pwd,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户密码
	Password *string `json:"password,omitempty"`
}

func (o SourceKafkaMqParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceKafkaMqParameters struct{}"
	}

	return strings.Join([]string{"SourceKafkaMqParameters", string(data)}, " ")
}

type SourceKafkaMqParametersSeekTo struct {
	value string
}

type SourceKafkaMqParametersSeekToEnum struct {
	LATEST   SourceKafkaMqParametersSeekTo
	EARLIEST SourceKafkaMqParametersSeekTo
}

func GetSourceKafkaMqParametersSeekToEnum() SourceKafkaMqParametersSeekToEnum {
	return SourceKafkaMqParametersSeekToEnum{
		LATEST: SourceKafkaMqParametersSeekTo{
			value: "latest",
		},
		EARLIEST: SourceKafkaMqParametersSeekTo{
			value: "earliest",
		},
	}
}

func (c SourceKafkaMqParametersSeekTo) Value() string {
	return c.value
}

func (c SourceKafkaMqParametersSeekTo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceKafkaMqParametersSeekTo) UnmarshalJSON(b []byte) error {
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

type SourceKafkaMqParametersSaslMechanism struct {
	value string
}

type SourceKafkaMqParametersSaslMechanismEnum struct {
	SCRAM_SHA_512 SourceKafkaMqParametersSaslMechanism
	PLAIN         SourceKafkaMqParametersSaslMechanism
}

func GetSourceKafkaMqParametersSaslMechanismEnum() SourceKafkaMqParametersSaslMechanismEnum {
	return SourceKafkaMqParametersSaslMechanismEnum{
		SCRAM_SHA_512: SourceKafkaMqParametersSaslMechanism{
			value: "SCRAM-SHA-512",
		},
		PLAIN: SourceKafkaMqParametersSaslMechanism{
			value: "PLAIN",
		},
	}
}

func (c SourceKafkaMqParametersSaslMechanism) Value() string {
	return c.value
}

func (c SourceKafkaMqParametersSaslMechanism) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceKafkaMqParametersSaslMechanism) UnmarshalJSON(b []byte) error {
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

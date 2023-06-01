package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	SeekTo *string `json:"seek_to,omitempty"`

	// SASL_SSL是否开启
	EnableSaslSsl *bool `json:"enable_sasl_ssl,omitempty"`

	// SASL认证机制
	SaslMechanism *string `json:"sasl_mechanism,omitempty"`

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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthInfo struct {

	// 用户安全集群的登录用户名
	AuthInfoName *string `json:"auth_info_name,omitempty"`

	// 用户安全集群的登录密码
	UserName *string `json:"user_name,omitempty"`

	// 用户安全集群的证书路径，目前只支持OBS路径，cer类型文件
	CertificateLocation *string `json:"certificate_location,omitempty"`

	// 数据源类型，目前支持CSS,KRB,passwd,Kafka_SSL
	DatasourceType *string `json:"datasource_type,omitempty"`

	// 创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// krb5配置文件obs路径
	Krb5Conf *string `json:"krb5_conf,omitempty"`

	// keytab配置文件obs路径
	Keytab *string `json:"keytab,omitempty"`

	// truststore配置文件obs路径
	TruststoreLocation *string `json:"truststore_location,omitempty"`

	// keystore配置文件obs路径
	KeystoreLocation *string `json:"keystore_location,omitempty"`

	// 所属用户名
	Owner *string `json:"owner,omitempty"`
}

func (o AuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthInfo struct{}"
	}

	return strings.Join([]string{"AuthInfo", string(data)}, " ")
}

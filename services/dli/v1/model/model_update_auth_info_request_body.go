package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAuthInfoRequestBody struct {

	// 证书名
	AuthInfoName string `json:"auth_info_name"`

	// 用户安全集群的新登录用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户安全集群的新登录密码
	Password *string `json:"password,omitempty"`

	// krb5配置文件obs路径
	Krb5Conf *string `json:"krb5_conf,omitempty"`

	// keytab配置文件obs路径
	Keytab *string `json:"keytab,omitempty"`

	// truststore配置文件obs路径
	TruststoreLocation *string `json:"truststore_location,omitempty"`

	// truststore配置文件密码
	TruststorePassword *string `json:"truststore_password,omitempty"`

	// keystore配置文件obs路径
	KeystoreLocation *string `json:"keystore_location,omitempty"`

	// keystore配置文件密码
	KeystorePassword *string `json:"keystore_password,omitempty"`
}

func (o UpdateAuthInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthInfoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAuthInfoRequestBody", string(data)}, " ")
}

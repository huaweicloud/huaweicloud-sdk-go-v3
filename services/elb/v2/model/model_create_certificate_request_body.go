package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateCertificateRequestBody struct {
	// 服务端公有密钥证书或者用于认证客户端证书的CA证书，由type字段区分。 格式：证书为PEM格式。

	Certificate string `json:"certificate"`
	// 服务端的私有密钥。 格式：私钥为PEM格式。 该字段仅type为server时有效且为必选。 该字段在type为client时无效。

	PrivateKey *string `json:"private_key,omitempty"`
	// SSL证书的描述信息。支持的最大字符长度：255

	Description *string `json:"description,omitempty"`
	// 服务端证书所签的域名。默认值：null；支持的最大字符长度：1024 取值范围： 普通域名由若干字符串组成，总长度为0-1024，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。 泛域名在普通域名的基础上仅允许首字母为\"*\"。 该字段仅type为server时有效。

	Domain *string `json:"domain,omitempty"`
	// SSL证书的名称。支持的最大字符长度：255

	Name *string `json:"name,omitempty"`
	// SSL证书的管理状态； 取值范围： true/false。 该字段为预留字段，暂未启用。只支持设定为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// SSL证书的类型。默认值：server； 取值范围： server：服务端证书； client：客户端证书；

	Type *string `json:"type,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}

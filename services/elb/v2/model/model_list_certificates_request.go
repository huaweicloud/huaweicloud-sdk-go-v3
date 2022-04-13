package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	// 每页返回的个数。 取值范围：0~intmax。

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询起始的证书id，为空时为查询第一页。 仅当和limit一起使用时生效

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 仅当和limit一起使用时生效。

	PageReverse *string `json:"page_reverse,omitempty"`
	// SSL证书ID。

	Id *string `json:"id,omitempty"`
	// SSL证书的名称。

	Name *string `json:"name,omitempty"`
	// 证书描述SSL证书描述。

	Description *string `json:"description,omitempty"`
	// SSL证书的类型。默认值：server；取值范围：server：服务端证书；client：客户端证书；

	Type *string `json:"type,omitempty"`
	// 服务端证书所签的域名。默认值：null；取值范围：普通域名由若干字符串组成，总长度为0-1024，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。泛域名在普通域名的基础上仅允许首字母为\"*\"。该字段仅type为server时有效。

	Domain *string `json:"domain,omitempty"`
	// PEM格式的服务端私有密钥。

	PrivateKey *string `json:"private_key,omitempty"`
	// PEM格式的服务端公有密钥或者用于认证客户端证书的CA证书，由type字段区分。

	Certificate *string `json:"certificate,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}

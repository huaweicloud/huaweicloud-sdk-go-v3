package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCertificateRequestBody This is a auto create Body Object
type UpdateCertificateRequestBody struct {

	// SSL证书对象 最大长度65536字符。 支持证书链，最大11层(含证书和证书链)
	Certificate *string `json:"certificate,omitempty"`

	// 服务端的私有密钥。  格式：私钥为PEM格式。 最大长度8192字符。
	PrivateKey *string `json:"private_key,omitempty"`

	// SSL证书的描述信息。  支持的最大字符长度：255
	Description *string `json:"description,omitempty"`

	// 服务端证书所签的域名。 取值：总长度为0-1024字符。  普通域名由若干字符串组成，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。  泛域名仅允许首段为\"*\"，其他取值限制与普通域名一致。如：*.domain.com，但不能是：*my.domain.com   该字段仅type为server时有效。
	Domain *string `json:"domain,omitempty"`

	// SSL证书的名称。  支持的最大字符长度：255
	Name *string `json:"name,omitempty"`

	// SSL证书的管理状态；  取值范围： true/false。  该字段为预留字段，暂未启用。只支持设定为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}

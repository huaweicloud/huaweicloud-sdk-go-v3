package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCredentialResponse Response Object
type ShowCredentialResponse struct {

	// AK。Access Key,是用来标识用户身份的访问密钥。
	Access *string `json:"access,omitempty"`

	// SK。Secret Key,用来对访问密钥进行加密签名,以验证身份。
	Secret *string `json:"secret,omitempty"`

	// security_token是将所获的AK、SK等信息进行加密后的字符串
	SecurityToken  *string `json:"security_token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCredentialResponse struct{}"
	}

	return strings.Join([]string{"ShowCredentialResponse", string(data)}, " ")
}

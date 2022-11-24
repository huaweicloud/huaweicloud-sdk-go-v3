package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCertKeyResponse struct {

	// 证书
	Cert *string `json:"cert,omitempty"`

	// 私钥
	PrivateKey     *string `json:"private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCertKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowCertKeyResponse", string(data)}, " ")
}

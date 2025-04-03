package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizationTokenResponse Response Object
type CreateAuthorizationTokenResponse struct {

	// 认证信息
	Auths map[string]AuthInfo `json:"auths,omitempty"`

	XSwrDockerlogin *string `json:"X-Swr-Dockerlogin,omitempty"`

	XSwrExpireat   *string `json:"x-swr-expireat,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuthorizationTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizationTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthorizationTokenResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretResponse Response Object
type CreateSecretResponse struct {

	// 认证信息
	Auths map[string]AuthInfo `json:"auths,omitempty"`

	XSwrDockerlogin *string `json:"X-Swr-Dockerlogin,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretResponse", string(data)}, " ")
}

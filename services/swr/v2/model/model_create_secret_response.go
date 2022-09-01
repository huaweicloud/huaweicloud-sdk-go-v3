package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSecretResponse struct {

	// 认证信息
	Auths map[string]AuthInfo `json:"auths,omitempty" xml:"auths"`

	XSwrDockerlogin *string `json:"X-Swr-Dockerlogin,omitempty" xml:"X-Swr-Dockerlogin"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretResponse", string(data)}, " ")
}

package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateKeypairResponse struct {
	// 密钥对名称。

	Name *string `json:"name,omitempty"`
	// 公钥。

	PublicKey *string `json:"public_key,omitempty"`
	// 私钥。

	PrivateKey *string `json:"private_key,omitempty"`
	// 用户ID。

	UserId *string `json:"user_id,omitempty"`
	// 指纹。

	Fingerprint    *string `json:"fingerprint,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeypairResponse struct{}"
	}

	return strings.Join([]string{"CreateKeypairResponse", string(data)}, " ")
}

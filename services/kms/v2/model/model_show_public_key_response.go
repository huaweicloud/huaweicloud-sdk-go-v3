package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPublicKeyResponse struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 公钥信息。
	PublicKey      *string `json:"public_key,omitempty" xml:"public_key"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPublicKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicKeyResponse", string(data)}, " ")
}

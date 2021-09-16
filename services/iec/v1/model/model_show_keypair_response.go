package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowKeypairResponse struct {
	// 密钥名称。

	Name *string `json:"name,omitempty"`
	//   密钥对应publicKey信息。

	PublicKey *string `json:"public_key,omitempty"`
	// 用户ID。

	UserId *string `json:"user_id,omitempty"`
	//   密钥对应指纹信息。

	Fingerprint    *string `json:"fingerprint,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKeypairResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKeypairResponse struct{}"
	}

	return strings.Join([]string{"ShowKeypairResponse", string(data)}, " ")
}

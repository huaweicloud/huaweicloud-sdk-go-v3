package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleKeypair 秘钥对对象
type SimpleKeypair struct {

	// 密钥名称。
	Name *string `json:"name,omitempty"`

	//   密钥对应publicKey信息。
	PublicKey *string `json:"public_key,omitempty"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	//   密钥对应指纹信息。
	Fingerprint *string `json:"fingerprint,omitempty"`
}

func (o SimpleKeypair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleKeypair struct{}"
	}

	return strings.Join([]string{"SimpleKeypair", string(data)}, " ")
}

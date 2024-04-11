package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerKeypair 服务器密钥对。
type ServerKeypair struct {

	// 云手机服务器ID，不得超过32个字节。
	ServerId string `json:"server_id"`

	// 密钥对名称。
	KeypairName string `json:"keypair_name"`
}

func (o ServerKeypair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerKeypair struct{}"
	}

	return strings.Join([]string{"ServerKeypair", string(data)}, " ")
}

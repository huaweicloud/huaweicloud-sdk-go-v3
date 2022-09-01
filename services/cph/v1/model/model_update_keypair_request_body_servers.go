package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateKeypairRequestBodyServers struct {

	// 密钥对名称
	KeypairName string `json:"keypair_name" xml:"keypair_name"`

	// 云手机服务器ID，不得超过32个字节
	ServerId string `json:"server_id" xml:"server_id"`
}

func (o UpdateKeypairRequestBodyServers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairRequestBodyServers struct{}"
	}

	return strings.Join([]string{"UpdateKeypairRequestBodyServers", string(data)}, " ")
}

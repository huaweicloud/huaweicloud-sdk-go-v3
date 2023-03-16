package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelHandshakeRequest struct {

	// 邀请的唯一标识符（ID）。帐号在发起邀请时创建ID。
	HandshakeId string `json:"handshake_id"`
}

func (o CancelHandshakeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelHandshakeRequest struct{}"
	}

	return strings.Join([]string{"CancelHandshakeRequest", string(data)}, " ")
}

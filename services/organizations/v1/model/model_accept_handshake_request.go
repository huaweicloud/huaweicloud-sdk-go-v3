package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptHandshakeRequest Request Object
type AcceptHandshakeRequest struct {

	// 邀请的唯一标识符（ID）。帐号在发起邀请时创建ID。
	HandshakeId string `json:"handshake_id"`
}

func (o AcceptHandshakeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptHandshakeRequest struct{}"
	}

	return strings.Join([]string{"AcceptHandshakeRequest", string(data)}, " ")
}

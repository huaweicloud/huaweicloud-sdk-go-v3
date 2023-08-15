package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHandshakeRequest Request Object
type ShowHandshakeRequest struct {

	// 邀请的唯一标识符（ID）。帐号在发起邀请时创建ID。
	HandshakeId string `json:"handshake_id"`
}

func (o ShowHandshakeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHandshakeRequest struct{}"
	}

	return strings.Join([]string{"ShowHandshakeRequest", string(data)}, " ")
}

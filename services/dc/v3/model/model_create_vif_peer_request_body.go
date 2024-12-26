package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVifPeerRequestBody 创建虚拟接口对等体请求体
type CreateVifPeerRequestBody struct {
	VifPeer *CreateVifPeer `json:"vif_peer,omitempty"`
}

func (o CreateVifPeerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVifPeerRequestBody", string(data)}, " ")
}

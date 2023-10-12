package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVifPeerRequestBody 更新对等体参数
type UpdateVifPeerRequestBody struct {
	VifPeer *UpdateVifPeer `json:"vif_peer,omitempty"`
}

func (o UpdateVifPeerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVifPeerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVifPeerRequestBody", string(data)}, " ")
}

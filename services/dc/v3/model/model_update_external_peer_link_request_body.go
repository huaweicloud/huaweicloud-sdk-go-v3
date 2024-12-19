package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExternalPeerLinkRequestBody 更新关联连接参数
type UpdateExternalPeerLinkRequestBody struct {
	PeerLink *UpdateExternalPeerLinkRequestBodyPeerLink `json:"peer_link,omitempty"`
}

func (o UpdateExternalPeerLinkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExternalPeerLinkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateExternalPeerLinkRequestBody", string(data)}, " ")
}

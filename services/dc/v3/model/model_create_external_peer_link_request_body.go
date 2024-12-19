package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalPeerLinkRequestBody 创建参数请求体
type CreateExternalPeerLinkRequestBody struct {
	PeerLink *CreateExternalPeerLinkRequestBodyPeerLink `json:"peer_link"`
}

func (o CreateExternalPeerLinkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalPeerLinkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateExternalPeerLinkRequestBody", string(data)}, " ")
}

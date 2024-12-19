package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPeerLinkResponse Response Object
type ShowPeerLinkResponse struct {
	PeerLink *PeerLinkEntry `json:"peer_link,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPeerLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPeerLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowPeerLinkResponse", string(data)}, " ")
}

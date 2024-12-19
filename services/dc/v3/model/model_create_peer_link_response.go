package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePeerLinkResponse Response Object
type CreatePeerLinkResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PeerLink       *ExternalCreatePeerLink `json:"peer_link,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreatePeerLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeerLinkResponse struct{}"
	}

	return strings.Join([]string{"CreatePeerLinkResponse", string(data)}, " ")
}

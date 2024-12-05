package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePeerLinkResponse Response Object
type UpdatePeerLinkResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PeerLink *ExternalUpdatePeerLink `json:"peer_link,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePeerLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeerLinkResponse struct{}"
	}

	return strings.Join([]string{"UpdatePeerLinkResponse", string(data)}, " ")
}

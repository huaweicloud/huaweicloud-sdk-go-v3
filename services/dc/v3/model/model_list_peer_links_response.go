package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPeerLinksResponse Response Object
type ListPeerLinksResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	PeerLinks *[]ExternalListPeerLinks `json:"peer_links,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPeerLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeerLinksResponse struct{}"
	}

	return strings.Join([]string{"ListPeerLinksResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPeerLinksResponse Response Object
type ListPeerLinksResponse struct {

	// 专线关联连接列表。
	PeerLinks *[]PeerLinkEntry `json:"peer_links,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPeerLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeerLinksResponse struct{}"
	}

	return strings.Join([]string{"ListPeerLinksResponse", string(data)}, " ")
}

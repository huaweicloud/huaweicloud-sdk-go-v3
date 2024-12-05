package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePeerLinkRequestBodyPeerLink struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	PeerSite *CreatePeerLinkRequestBodyPeerLinkPeerSite `json:"peer_site"`
}

func (o CreatePeerLinkRequestBodyPeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeerLinkRequestBodyPeerLink struct{}"
	}

	return strings.Join([]string{"CreatePeerLinkRequestBodyPeerLink", string(data)}, " ")
}

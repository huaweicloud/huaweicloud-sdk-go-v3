package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePeerLinkRequestBody struct {

	// 空运行 - true 是 - false 否
	DryRun *bool `json:"dry_run,omitempty"`

	PeerLink *UpdatePeerLinkRequestBodyPeerLink `json:"peer_link,omitempty"`
}

func (o UpdatePeerLinkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeerLinkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePeerLinkRequestBody", string(data)}, " ")
}

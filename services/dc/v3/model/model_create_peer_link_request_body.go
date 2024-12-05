package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePeerLinkRequestBody struct {

	// 空运行 - true 是 - false 否
	DryRun *bool `json:"dry_run,omitempty"`

	PeerLink *CreatePeerLinkRequestBodyPeerLink `json:"peer_link,omitempty"`
}

func (o CreatePeerLinkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeerLinkRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePeerLinkRequestBody", string(data)}, " ")
}

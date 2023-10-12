package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVifPeerRequestBody struct {
	VifPeer *CreateVifPeer `json:"vif_peer,omitempty"`
}

func (o CreateVifPeerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVifPeerRequestBody", string(data)}, " ")
}

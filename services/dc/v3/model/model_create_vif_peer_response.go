package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVifPeerResponse Response Object
type CreateVifPeerResponse struct {
	VifPeer        *VifPeer `json:"vif_peer,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CreateVifPeerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerResponse struct{}"
	}

	return strings.Join([]string{"CreateVifPeerResponse", string(data)}, " ")
}

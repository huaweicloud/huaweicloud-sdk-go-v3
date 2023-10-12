package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVifPeerResponse Response Object
type UpdateVifPeerResponse struct {
	VifPeer        *VifPeer `json:"vif_peer,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateVifPeerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVifPeerResponse struct{}"
	}

	return strings.Join([]string{"UpdateVifPeerResponse", string(data)}, " ")
}

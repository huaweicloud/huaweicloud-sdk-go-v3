package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVifPeerRequest Request Object
type UpdateVifPeerRequest struct {

	// 虚拟接口对等体ID
	VifPeerId string `json:"vif_peer_id"`

	Body *UpdateVifPeerRequestBody `json:"body,omitempty"`
}

func (o UpdateVifPeerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVifPeerRequest struct{}"
	}

	return strings.Join([]string{"UpdateVifPeerRequest", string(data)}, " ")
}

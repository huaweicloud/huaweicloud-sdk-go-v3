package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVifPeerRequest Request Object
type DeleteVifPeerRequest struct {

	// 虚拟接口对等体ID
	VifPeerId string `json:"vif_peer_id"`
}

func (o DeleteVifPeerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVifPeerRequest struct{}"
	}

	return strings.Join([]string{"DeleteVifPeerRequest", string(data)}, " ")
}

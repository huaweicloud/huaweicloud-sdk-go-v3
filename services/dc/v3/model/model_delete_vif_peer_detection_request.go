package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVifPeerDetectionRequest Request Object
type DeleteVifPeerDetectionRequest struct {

	// 虚拟接口对等体ID
	VifPeerId string `json:"vif_peer_id"`
}

func (o DeleteVifPeerDetectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVifPeerDetectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteVifPeerDetectionRequest", string(data)}, " ")
}

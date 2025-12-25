package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVifPeerDetectionRequest Request Object
type ShowVifPeerDetectionRequest struct {

	// 虚拟接口对等体ID
	VifPeerId string `json:"vif_peer_id"`
}

func (o ShowVifPeerDetectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVifPeerDetectionRequest struct{}"
	}

	return strings.Join([]string{"ShowVifPeerDetectionRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVifPeerDetectionRequestBodyVifPeerDetection struct {

	// 虚拟接口对等体实例id
	VifPeerId string `json:"vif_peer_id"`
}

func (o CreateVifPeerDetectionRequestBodyVifPeerDetection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerDetectionRequestBodyVifPeerDetection struct{}"
	}

	return strings.Join([]string{"CreateVifPeerDetectionRequestBodyVifPeerDetection", string(data)}, " ")
}

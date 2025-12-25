package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVifPeerDetectionResponse Response Object
type CreateVifPeerDetectionResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	VifPeerDetection *VifPeerDetection `json:"vif_peer_detection,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateVifPeerDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerDetectionResponse struct{}"
	}

	return strings.Join([]string{"CreateVifPeerDetectionResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVifPeerDetectionResponse Response Object
type ShowVifPeerDetectionResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	VifPeerDetection *VifPeerDetection `json:"vif_peer_detection,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowVifPeerDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVifPeerDetectionResponse struct{}"
	}

	return strings.Join([]string{"ShowVifPeerDetectionResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVifPeerDetectionsResponse Response Object
type ListVifPeerDetectionsResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	VifPeerDetections *[]VifPeerDetection `json:"vif_peer_detections,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListVifPeerDetectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVifPeerDetectionsResponse struct{}"
	}

	return strings.Join([]string{"ListVifPeerDetectionsResponse", string(data)}, " ")
}

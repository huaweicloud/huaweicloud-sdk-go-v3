package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVifPeerDetectionRequestBody 创建虚拟接口对等体连通性探测实例请求体
type CreateVifPeerDetectionRequestBody struct {
	VifPeerDetection *CreateVifPeerDetectionRequestBodyVifPeerDetection `json:"vif_peer_detection,omitempty"`
}

func (o CreateVifPeerDetectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerDetectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVifPeerDetectionRequestBody", string(data)}, " ")
}

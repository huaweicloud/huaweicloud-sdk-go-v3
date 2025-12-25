package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVifPeerDetectionRequest Request Object
type CreateVifPeerDetectionRequest struct {
	Body *CreateVifPeerDetectionRequestBody `json:"body,omitempty"`
}

func (o CreateVifPeerDetectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerDetectionRequest struct{}"
	}

	return strings.Join([]string{"CreateVifPeerDetectionRequest", string(data)}, " ")
}

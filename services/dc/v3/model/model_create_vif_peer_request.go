package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVifPeerRequest Request Object
type CreateVifPeerRequest struct {
	Body *CreateVifPeerRequestBody `json:"body,omitempty"`
}

func (o CreateVifPeerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVifPeerRequest struct{}"
	}

	return strings.Join([]string{"CreateVifPeerRequest", string(data)}, " ")
}

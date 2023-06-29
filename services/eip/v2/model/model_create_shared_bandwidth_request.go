package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSharedBandwidthRequest Request Object
type CreateSharedBandwidthRequest struct {
	Body *CreateSharedBandwidhRequestBody `json:"body,omitempty"`
}

func (o CreateSharedBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedBandwidthRequest struct{}"
	}

	return strings.Join([]string{"CreateSharedBandwidthRequest", string(data)}, " ")
}

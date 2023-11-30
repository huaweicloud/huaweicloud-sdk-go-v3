package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyBandwidthRequest Request Object
type BatchModifyBandwidthRequest struct {
	Body *ModifyBandwidthRequestBody `json:"body,omitempty"`
}

func (o BatchModifyBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyBandwidthRequest struct{}"
	}

	return strings.Join([]string{"BatchModifyBandwidthRequest", string(data)}, " ")
}

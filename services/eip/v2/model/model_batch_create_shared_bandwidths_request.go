package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateSharedBandwidthsRequest struct {
	Body *BatchCreateBandwidthRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchCreateSharedBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSharedBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSharedBandwidthsRequest", string(data)}, " ")
}

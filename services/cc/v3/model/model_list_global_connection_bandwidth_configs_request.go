package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConnectionBandwidthConfigsRequest Request Object
type ListGlobalConnectionBandwidthConfigsRequest struct {
}

func (o ListGlobalConnectionBandwidthConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthConfigsRequest", string(data)}, " ")
}

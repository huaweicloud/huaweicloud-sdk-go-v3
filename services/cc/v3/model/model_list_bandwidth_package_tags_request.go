package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageTagsRequest Request Object
type ListBandwidthPackageTagsRequest struct {
}

func (o ListBandwidthPackageTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageTagsRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageTagsRequest", string(data)}, " ")
}

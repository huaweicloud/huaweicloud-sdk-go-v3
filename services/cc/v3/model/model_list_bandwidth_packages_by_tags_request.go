package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackagesByTagsRequest Request Object
type ListBandwidthPackagesByTagsRequest struct {
	Body *ListBandwidthPackagesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListBandwidthPackagesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackagesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackagesByTagsRequest", string(data)}, " ")
}

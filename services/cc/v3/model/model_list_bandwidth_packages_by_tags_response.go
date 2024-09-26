package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackagesByTagsResponse Response Object
type ListBandwidthPackagesByTagsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 带宽包实例列表。
	BandwidthPackages []BandwidthPackage `json:"bandwidth_packages"`
	HttpStatusCode    int                `json:"-"`
}

func (o ListBandwidthPackagesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackagesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackagesByTagsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegionBandwidthPackageSpecificationsResponse Response Object
type ListRegionBandwidthPackageSpecificationsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 租户带宽包城域规格列表。
	RegionSpecifications []RegionBandwidthPackageSpecification `json:"region_specifications"`
	HttpStatusCode       int                                   `json:"-"`
}

func (o ListRegionBandwidthPackageSpecificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionBandwidthPackageSpecificationsResponse struct{}"
	}

	return strings.Join([]string{"ListRegionBandwidthPackageSpecificationsResponse", string(data)}, " ")
}

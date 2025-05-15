package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegionBandwidthPackageSpecificationsRequest Request Object
type ListRegionBandwidthPackageSpecificationsRequest struct {

	// 根据城域带宽包本端区域ID过滤租户城域带宽配置列表
	LocalRegionId *string `json:"local_region_id,omitempty"`

	// 根据城域带宽包对端区域ID过滤租户城域带宽配置列表
	RemoteRegionId *string `json:"remote_region_id,omitempty"`
}

func (o ListRegionBandwidthPackageSpecificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionBandwidthPackageSpecificationsRequest struct{}"
	}

	return strings.Join([]string{"ListRegionBandwidthPackageSpecificationsRequest", string(data)}, " ")
}

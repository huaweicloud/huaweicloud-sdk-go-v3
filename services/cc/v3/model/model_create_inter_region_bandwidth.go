package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInterRegionBandwidth 创建域间带宽的详情信息。
type CreateInterRegionBandwidth struct {

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 带宽包实例ID。
	BandwidthPackageId string `json:"bandwidth_package_id"`

	// 域间带宽值。
	Bandwidth int32 `json:"bandwidth"`

	// 域间RegionID。
	InterRegionIds []string `json:"inter_region_ids"`
}

func (o CreateInterRegionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInterRegionBandwidth struct{}"
	}

	return strings.Join([]string{"CreateInterRegionBandwidth", string(data)}, " ")
}

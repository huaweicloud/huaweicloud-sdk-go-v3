package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkCapabilityEntry 分支网络租户能力详情条目。
type SiteNetworkCapabilityEntry struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	Specification *SiteNetworkSpecificationEnum `json:"specification"`

	// 是否支持分支网络。
	IsSupport *bool `json:"is_support,omitempty"`

	// 是否支持分支网络企业项目。
	IsSupportEnterpriseProject *bool `json:"is_support_enterprise_project,omitempty"`

	// 是否支持分支网络标签。
	IsSupportTag *bool `json:"is_support_tag,omitempty"`

	// 是否支持创建同region分支网络。
	IsSupportIntraRegion *bool `json:"is_support_intra_region,omitempty"`

	// 分支网络的拓扑列表。
	SupportTopologies *[]SiteNetworkTopologyEnum `json:"support_topologies,omitempty"`

	// list类型
	SupportRegions *[]string `json:"support_regions,omitempty"`

	// list类型
	SupportDscpRegions *[]string `json:"support_dscp_regions,omitempty"`

	// list类型
	SupportFreezeRegions *[]string `json:"support_freeze_regions,omitempty"`

	// list类型
	SupportLocations *[]string `json:"support_locations,omitempty"`

	SizeRange *ConnectionBandwidthSizeRange `json:"size_range,omitempty"`

	// list类型
	ChargeMode *[]ConnectionBandwidthChargeModeEnum `json:"charge_mode,omitempty"`
}

func (o SiteNetworkCapabilityEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkCapabilityEntry struct{}"
	}

	return strings.Join([]string{"SiteNetworkCapabilityEntry", string(data)}, " ")
}

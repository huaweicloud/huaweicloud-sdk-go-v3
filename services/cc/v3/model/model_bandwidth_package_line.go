package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthPackageLine 带宽包线路。
type BandwidthPackageLine struct {

	// RegionID。
	LocalRegionId *string `json:"local_region_id,omitempty"`

	// RegionID。
	RemoteRegionId *string `json:"remote_region_id,omitempty"`

	// 站点编码。
	LocalSiteCode *string `json:"local_site_code,omitempty"`

	// 站点编码。
	RemoteSiteCode *string `json:"remote_site_code,omitempty"`

	// 支持的等级列表。
	SupportLevels *[]string `json:"support_levels,omitempty"`

	// 产品编码列表。
	SpecCodes *[]BandwidthPackageLineSpecCode `json:"spec_codes,omitempty"`
}

func (o BandwidthPackageLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackageLine struct{}"
	}

	return strings.Join([]string{"BandwidthPackageLine", string(data)}, " ")
}

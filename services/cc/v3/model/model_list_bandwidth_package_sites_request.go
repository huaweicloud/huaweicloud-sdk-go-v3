package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageSitesRequest Request Object
type ListBandwidthPackageSitesRequest struct {

	// 根据站点编码进行查询
	SiteCode *string `json:"site_code,omitempty"`

	// 根据区域ID进行查询
	RegionId *string `json:"region_id,omitempty"`

	// 根据名称模糊查询
	Name *string `json:"name,omitempty"`
}

func (o ListBandwidthPackageSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageSitesRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageSitesRequest", string(data)}, " ")
}

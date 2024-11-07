package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageSitesResponse Response Object
type ListBandwidthPackageSitesResponse struct {

	// 站点列表表。
	BandwidthPackageSites *[]BandwidthPackageSite `json:"bandwidth_package_sites,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBandwidthPackageSitesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageSitesResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageSitesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBandwidthPkgResponse struct {

	// bandwidthpkg对象列表
	Bandwidthpkgs *[]BandwidthPkgResp `json:"bandwidthpkgs,omitempty"`

	// 翻页展示
	BandwidthpkgsLinks *[]BandwidthPkgPage `json:"bandwidthpkgs_links,omitempty"`
	HttpStatusCode     int                 `json:"-"`
}

func (o ListBandwidthPkgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPkgResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPkgResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageLevelsResponse Response Object
type ListBandwidthPackageLevelsResponse struct {

	// 带宽包等级列表。
	BandwidthPackageLevels *[]BandwidthPackageLevel `json:"bandwidth_package_levels,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBandwidthPackageLevelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageLevelsResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageLevelsResponse", string(data)}, " ")
}

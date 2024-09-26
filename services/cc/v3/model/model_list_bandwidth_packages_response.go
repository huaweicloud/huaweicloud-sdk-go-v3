package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackagesResponse Response Object
type ListBandwidthPackagesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 带宽包实例列表。
	BandwidthPackages []BandwidthPackage `json:"bandwidth_packages"`
	HttpStatusCode    int                `json:"-"`
}

func (o ListBandwidthPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackagesResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackagesResponse", string(data)}, " ")
}

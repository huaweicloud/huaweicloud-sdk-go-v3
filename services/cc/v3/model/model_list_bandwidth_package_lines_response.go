package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageLinesResponse Response Object
type ListBandwidthPackageLinesResponse struct {

	// 带宽包线路列表。
	BandwidthPackageLines *[]BandwidthPackageLine `json:"bandwidth_package_lines,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBandwidthPackageLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageLinesResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageLinesResponse", string(data)}, " ")
}

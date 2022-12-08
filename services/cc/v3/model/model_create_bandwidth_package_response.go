package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateBandwidthPackageResponse struct {
	BandwidthPackage *BandwidthPackage `json:"bandwidth_package,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"CreateBandwidthPackageResponse", string(data)}, " ")
}

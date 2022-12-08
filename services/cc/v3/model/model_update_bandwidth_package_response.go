package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateBandwidthPackageResponse struct {
	BandwidthPackage *BandwidthPackage `json:"bandwidth_package,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthPackageResponse", string(data)}, " ")
}

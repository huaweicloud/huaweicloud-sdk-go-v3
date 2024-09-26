package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBandwidthPackageResponse Response Object
type CreateBandwidthPackageResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	BandwidthPackage *BandwidthPackage `json:"bandwidth_package"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"CreateBandwidthPackageResponse", string(data)}, " ")
}

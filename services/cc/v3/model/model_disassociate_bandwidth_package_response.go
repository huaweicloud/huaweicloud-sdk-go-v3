package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateBandwidthPackageResponse Response Object
type DisassociateBandwidthPackageResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	BandwidthPackage *BandwidthPackage `json:"bandwidth_package"`
	HttpStatusCode   int               `json:"-"`
}

func (o DisassociateBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"DisassociateBandwidthPackageResponse", string(data)}, " ")
}

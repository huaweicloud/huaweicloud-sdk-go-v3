package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateBandwidthPackageResponse Response Object
type AssociateBandwidthPackageResponse struct {
	BandwidthPackage *BandwidthPackage `json:"bandwidth_package,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"AssociateBandwidthPackageResponse", string(data)}, " ")
}

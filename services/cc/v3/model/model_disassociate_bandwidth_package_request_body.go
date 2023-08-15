package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateBandwidthPackageRequestBody 解关联带宽包实例的请求体。
type DisassociateBandwidthPackageRequestBody struct {
	BandwidthPackage *DisassociateBandwidthPackage `json:"bandwidth_package"`
}

func (o DisassociateBandwidthPackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateBandwidthPackageRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateBandwidthPackageRequestBody", string(data)}, " ")
}

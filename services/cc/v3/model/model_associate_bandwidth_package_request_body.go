package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateBandwidthPackageRequestBody 关联带宽包实例的请求体。
type AssociateBandwidthPackageRequestBody struct {
	BandwidthPackage *AssociateBandwidthPackage `json:"bandwidth_package"`
}

func (o AssociateBandwidthPackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateBandwidthPackageRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateBandwidthPackageRequestBody", string(data)}, " ")
}

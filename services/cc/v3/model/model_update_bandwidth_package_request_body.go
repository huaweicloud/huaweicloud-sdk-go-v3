package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新带宽包实例的请求体。
type UpdateBandwidthPackageRequestBody struct {
	BandwidthPackage *UpdateBandwidthPackage `json:"bandwidth_package"`
}

func (o UpdateBandwidthPackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthPackageRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthPackageRequestBody", string(data)}, " ")
}

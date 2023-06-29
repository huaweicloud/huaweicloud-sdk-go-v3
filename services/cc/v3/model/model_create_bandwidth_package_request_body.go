package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBandwidthPackageRequestBody 创建带宽包实例的请求体。
type CreateBandwidthPackageRequestBody struct {
	BandwidthPackage *CreateBandwidthPackage `json:"bandwidth_package"`
}

func (o CreateBandwidthPackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBandwidthPackageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBandwidthPackageRequestBody", string(data)}, " ")
}

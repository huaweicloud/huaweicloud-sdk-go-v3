package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBandwidthPackageRequest Request Object
type DeleteBandwidthPackageRequest struct {

	// 带宽包实例ID。
	Id string `json:"id"`
}

func (o DeleteBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"DeleteBandwidthPackageRequest", string(data)}, " ")
}

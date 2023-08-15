package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthPackageRequest Request Object
type ShowBandwidthPackageRequest struct {

	// 带宽包实例ID。
	Id string `json:"id"`
}

func (o ShowBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthPackageRequest", string(data)}, " ")
}

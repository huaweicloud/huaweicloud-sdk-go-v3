package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagBandwidthPackageRequest Request Object
type UntagBandwidthPackageRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *UntagBandwidthPackageRequestBody `json:"body,omitempty"`
}

func (o UntagBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"UntagBandwidthPackageRequest", string(data)}, " ")
}

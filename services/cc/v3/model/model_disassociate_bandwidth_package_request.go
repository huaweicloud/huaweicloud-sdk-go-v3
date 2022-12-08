package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateBandwidthPackageRequest struct {

	// 带宽包实例ID。
	Id string `json:"id"`

	Body *DisassociateBandwidthPackageRequestBody `json:"body,omitempty"`
}

func (o DisassociateBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"DisassociateBandwidthPackageRequest", string(data)}, " ")
}

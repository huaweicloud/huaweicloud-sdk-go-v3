package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateBandwidthPackageRequest Request Object
type AssociateBandwidthPackageRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *AssociateBandwidthPackageRequestBody `json:"body,omitempty"`
}

func (o AssociateBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"AssociateBandwidthPackageRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagBandwidthPackageRequest Request Object
type TagBandwidthPackageRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *TagBandwidthPackageRequestBody `json:"body,omitempty"`
}

func (o TagBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"TagBandwidthPackageRequest", string(data)}, " ")
}

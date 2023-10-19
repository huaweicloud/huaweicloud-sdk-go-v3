package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagBandwidthPackageRequestBody 删除带宽包标签的请求体。
type UntagBandwidthPackageRequestBody struct {

	// 包含标签。
	Tags []Tag `json:"tags"`
}

func (o UntagBandwidthPackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagBandwidthPackageRequestBody struct{}"
	}

	return strings.Join([]string{"UntagBandwidthPackageRequestBody", string(data)}, " ")
}

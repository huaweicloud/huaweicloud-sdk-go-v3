package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagBandwidthPackageRequestBody 创建带宽包标签的请求体。
type TagBandwidthPackageRequestBody struct {

	// 包含标签。
	Tags []Tag `json:"tags"`
}

func (o TagBandwidthPackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagBandwidthPackageRequestBody struct{}"
	}

	return strings.Join([]string{"TagBandwidthPackageRequestBody", string(data)}, " ")
}

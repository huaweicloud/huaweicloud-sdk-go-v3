package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackagesByTagsRequestBody 通过标签过滤带宽包实例的请求体。
type ListBandwidthPackagesByTagsRequestBody struct {

	// 包含标签。
	Tags []MultivaluedTag `json:"tags"`
}

func (o ListBandwidthPackagesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackagesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackagesByTagsRequestBody", string(data)}, " ")
}

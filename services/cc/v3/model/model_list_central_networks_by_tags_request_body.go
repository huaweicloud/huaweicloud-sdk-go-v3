package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworksByTagsRequestBody 通过标签过滤中心网络实例的请求体。
type ListCentralNetworksByTagsRequestBody struct {

	// 包含标签。
	Tags []MultivaluedTag `json:"tags"`
}

func (o ListCentralNetworksByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworksByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListCentralNetworksByTagsRequestBody", string(data)}, " ")
}

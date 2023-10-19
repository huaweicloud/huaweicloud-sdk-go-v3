package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionsByTagsRequestBody 通过标签过滤云连接实例的请求体。
type ListCloudConnectionsByTagsRequestBody struct {

	// 包含标签。
	Tags []MultivaluedTag `json:"tags"`
}

func (o ListCloudConnectionsByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionsByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionsByTagsRequestBody", string(data)}, " ")
}

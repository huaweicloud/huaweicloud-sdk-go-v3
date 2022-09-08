package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQueryProjectResourceTagsResponse struct {

	// 包含标签，最多包含10个key，每 个key下面的value最多10个，每 个key对应的value可以为空数组但 结构体不能缺失。Key不能重复， 同一个key中values不能重复。结 果返回包含所有标签的资源列表， key之间是与的关系，key-value结 构中value是或的关系。无tag过滤 条件时返回全量数据。
	Tags           *[]TagValuesList `json:"tags,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryProjectResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryProjectResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryProjectResourceTagsResponse", string(data)}, " ")
}

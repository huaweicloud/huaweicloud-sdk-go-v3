package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListResourceInstancesRequestBody struct {

	// 包含标签，最多包含50个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]TagsFilter `json:"tags,omitempty"`
}

func (o ListResourceInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTagValuesRequest struct {

	// 存储ID
	DataStoreId string `json:"data_store_id"`

	// tag 名称
	TagName string `json:"tag_name"`

	// 查询标签的值的过滤条件，例如: {\"deviceCategory\": \"class1\"}，注意特殊字符需要 urlencode
	Filters *string `json:"filters,omitempty"`

	// 查询起始元素的偏移
	Offset *int32 `json:"offset,omitempty"`

	// 返回的元素列表大小限制,默认为 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTagValuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagValuesRequest struct{}"
	}

	return strings.Join([]string{"ListTagValuesRequest", string(data)}, " ")
}

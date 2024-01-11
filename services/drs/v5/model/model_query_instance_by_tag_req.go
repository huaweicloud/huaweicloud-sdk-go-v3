package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryInstanceByTagReq 查询资源实例请求
type QueryInstanceByTagReq struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”字段。该字段为false或者未提供该参数时，该条件不生效，即返回所有资源或按\"tags\"，\"matches\"等条件过滤。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含20个key，每个key下面的value最多20个。
	Tags *[]Tag `json:"tags,omitempty"`

	// 搜索字段,key为要匹配的字段，如resource_name等。value为匹配的值。
	Matches *[]TagMatch `json:"matches,omitempty"`
}

func (o QueryInstanceByTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInstanceByTagReq struct{}"
	}

	return strings.Join([]string{"QueryInstanceByTagReq", string(data)}, " ")
}

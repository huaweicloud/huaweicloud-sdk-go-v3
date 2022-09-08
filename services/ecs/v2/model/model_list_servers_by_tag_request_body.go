package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ListServersByTagRequestBody struct {

	// 值为filter：表示按标签过滤弹性云服务器，返回符合条件的云服务器列表。
	Action string `json:"action"`

	// 查询返回的云服务器数量限制，最多为1000，不能为负数。  - 如果action的值为count时，此参数无效。 - 如果action的值为filter时，limit必填，取值范围[0-1000]，如果不传值，系统默认limit值为1000。
	Limit *string `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0。  查询第一页数据时，可以不传入此参数。  - 如果action的值为count，此参数无效。 - 如果action的值为filter时，必填，如果用户不传值，系统默认offset值为0。
	Offset *string `json:"offset,omitempty"`

	// 查询包含所有指定标签的弹性云服务器。  - 最多包含10个key，每个key下面的value最多10个。 - 结构体不能缺失。 - key不能为空或者空字符串。 - key不能重复。 - 同一个key中values不能重复。
	Tags *[]ServerTags `json:"tags,omitempty"`

	// 查询不包含所有指定标签的弹性云服务器。  - 最多包含10个key，每个key下面的value最多10个。 - 结构体不能缺失。 - key不能为空或者空字符串。 - key不能重复。 - 同一个key中values不能重复。
	NotTags *[]ServerTags `json:"not_tags,omitempty"`

	// 搜索字段，用于按条件搜索弹性云服务器。  当前仅支持按resource_name进行搜索
	Matches *[]ServerTagMatch `json:"matches,omitempty"`
}

func (o ListServersByTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersByTagRequestBody struct{}"
	}

	return strings.Join([]string{"ListServersByTagRequestBody", string(data)}, " ")
}

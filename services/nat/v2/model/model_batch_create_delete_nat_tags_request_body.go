package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteNatTagsRequestBody 请求参数。
type BatchCreateDeleteNatTagsRequestBody struct {

	// 标签列表。请参考表Tags字段数据结构说明
	Tags []PublicTags `json:"tags"`

	// 操作标识：仅限于create（创建）、delete（删除）
	Action string `json:"action"`
}

func (o BatchCreateDeleteNatTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteNatTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteNatTagsRequestBody", string(data)}, " ")
}

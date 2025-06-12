package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagsRequestBody 批量删除资源标签请求体。
type BatchDeleteResourceTagsRequestBody struct {

	// 标签列表
	Tags []Tag `json:"tags"`
}

func (o BatchDeleteResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagsRequestBody", string(data)}, " ")
}

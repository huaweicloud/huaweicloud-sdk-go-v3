package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourceTagsRequestBody 批量添加资源标签请求体。
type BatchCreateResourceTagsRequestBody struct {

	// 标签列表
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o BatchCreateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsRequestBody", string(data)}, " ")
}

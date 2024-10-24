package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTagsRequestBody 批量打资源标签请求体
type BatchCreateTagsRequestBody struct {

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 系统标签列表。
	SysTags *[]SystemTag `json:"sys_tags,omitempty"`
}

func (o BatchCreateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsRequestBody", string(data)}, " ")
}

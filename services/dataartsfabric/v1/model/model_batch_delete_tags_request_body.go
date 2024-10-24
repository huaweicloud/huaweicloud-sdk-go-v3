package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagsRequestBody 批量删除的资源标签
type BatchDeleteTagsRequestBody struct {

	// 批量删除请求体标签列表。
	Tags *[]DeleteResourceTag `json:"tags,omitempty"`

	// 批量删除请求体系统标签列表。
	SysTags *[]DeleteResourceTag `json:"sys_tags,omitempty"`
}

func (o BatchDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTagsRequestBody struct {

	// 标签对象列表
	Tags *[]ResourceTagInfo `json:"tags,omitempty"`

	// 系统标签对象列表
	SysTags *[]ResourceTagInfo `json:"sys_tags,omitempty"`
}

func (o BatchDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequestBody", string(data)}, " ")
}

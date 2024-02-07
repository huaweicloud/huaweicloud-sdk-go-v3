package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateV2RequestBody struct {

	// 全域弹性公网IP标签
	Tags []BatchCreateV2RequestBodyTags `json:"tags"`

	// 系统标签
	SysTags *[]BatchCreateV2RequestBodySysTags `json:"sys_tags,omitempty"`
}

func (o BatchCreateV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateV2RequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateV2RequestBody", string(data)}, " ")
}

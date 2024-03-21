package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteV2RequestBody struct {

	// 全域弹性公网IP标签
	Tags []BatchDeleteV2RequestBodyTags `json:"tags"`

	// 系统标签
	SysTags *[]BatchDeleteV2RequestBodyTags `json:"sys_tags,omitempty"`
}

func (o BatchDeleteV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteV2RequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteV2RequestBody", string(data)}, " ")
}

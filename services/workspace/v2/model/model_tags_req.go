package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsReq 批量修改标签请求
type TagsReq struct {

	// 操作字符串create或delete
	Action *string `json:"action,omitempty"`

	// 标签对象
	Tags *[]Tag `json:"tags,omitempty"`

	// 标签对象
	SysTags *[]Tag `json:"sys_tags,omitempty"`
}

func (o TagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsReq struct{}"
	}

	return strings.Join([]string{"TagsReq", string(data)}, " ")
}

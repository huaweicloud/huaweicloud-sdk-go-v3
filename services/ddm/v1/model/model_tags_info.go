package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsInfo struct {

	// 标签。
	Tags *[]Tags `json:"tags,omitempty"`

	// 系统标签。
	SysTags *[]Tags `json:"sys_tags,omitempty"`
}

func (o TagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsInfo struct{}"
	}

	return strings.Join([]string{"TagsInfo", string(data)}, " ")
}

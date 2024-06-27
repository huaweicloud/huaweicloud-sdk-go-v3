package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplatePageParam struct {

	// 创建人ID
	CreatorNum *string `json:"creator_num,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	// 脑图名称
	Name *string `json:"name,omitempty"`
}

func (o TemplatePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplatePageParam struct{}"
	}

	return strings.Join([]string{"TemplatePageParam", string(data)}, " ")
}

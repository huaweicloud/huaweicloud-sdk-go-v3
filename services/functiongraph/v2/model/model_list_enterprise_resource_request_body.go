package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEnterpriseResourceRequestBody struct {

	// 是否通过标签过滤
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 每页显示条数
	Limit *string `json:"limit,omitempty"`

	// 查询偏移量
	Offset *string `json:"offset,omitempty"`

	// 查询指定action
	Action *string `json:"action,omitempty"`

	// 查询指定键值对
	Matches *[]KvItem `json:"matches,omitempty"`

	// 查询指定系统标签列表
	SysTags *[]TagItem `json:"sys_tags,omitempty"`
}

func (o ListEnterpriseResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseResourceRequestBody struct{}"
	}

	return strings.Join([]string{"ListEnterpriseResourceRequestBody", string(data)}, " ")
}

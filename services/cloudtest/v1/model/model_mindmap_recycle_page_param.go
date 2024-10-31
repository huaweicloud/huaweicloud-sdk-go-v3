package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MindmapRecyclePageParam struct {

	// 创建人工号
	CreatorNum *string `json:"creator_num,omitempty"`

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 过滤用字段
	Text *string `json:"text,omitempty"`
}

func (o MindmapRecyclePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindmapRecyclePageParam struct{}"
	}

	return strings.Join([]string{"MindmapRecyclePageParam", string(data)}, " ")
}

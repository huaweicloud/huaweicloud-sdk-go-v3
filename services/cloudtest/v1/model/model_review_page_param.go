package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReviewPageParam struct {
	Deleted *string `json:"deleted,omitempty"`

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	MindmapId *string `json:"mindmap_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeValue *string `json:"node_value,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o ReviewPageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewPageParam struct{}"
	}

	return strings.Join([]string{"ReviewPageParam", string(data)}, " ")
}

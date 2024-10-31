package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePageInfoMindmapRecycle struct {

	// 总条数
	Total *int64 `json:"total,omitempty"`

	// 入参集合
	List *[]MindmapRecycle `json:"list,omitempty"`

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	// 总页数
	Pages *int32 `json:"pages,omitempty"`

	// 每页条数
	Size *int32 `json:"size,omitempty"`
}

func (o BasePageInfoMindmapRecycle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePageInfoMindmapRecycle struct{}"
	}

	return strings.Join([]string{"BasePageInfoMindmapRecycle", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotesRequestBody 查询评论列表请求体
type ListNotesRequestBody struct {

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询参数。用于指定查询结果的起始位置，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段
	SortBy *string `json:"sort_by,omitempty"`

	// 升序/降序
	Order *string `json:"order,omitempty"`

	// 搜索起始时间
	FromDate *string `json:"from_date,omitempty"`

	// 搜索结束时间
	ToDate *string `json:"to_date,omitempty"`

	// 评论的对象ID
	WarRoomId *string `json:"war_room_id,omitempty"`
}

func (o ListNotesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotesRequestBody struct{}"
	}

	return strings.Join([]string{"ListNotesRequestBody", string(data)}, " ")
}

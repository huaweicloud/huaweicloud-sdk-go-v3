package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLimitTaskResponse Response Object
type ListLimitTaskResponse struct {

	// 限流任务列表
	LimitTaskList *[]ListLimitTaskResponseResult `json:"limit_task_list,omitempty"`

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置。
	Offset *int32 `json:"offset,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"ListLimitTaskResponse", string(data)}, " ")
}

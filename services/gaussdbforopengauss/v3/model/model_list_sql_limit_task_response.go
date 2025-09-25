package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlLimitTaskResponse Response Object
type ListSqlLimitTaskResponse struct {

	// **参数解释**: 限流任务列表。
	LimitTaskList *[]ListSqlLimitTaskResponseResult `json:"limit_task_list,omitempty"`

	// **参数解释**: 查询记录数。 **取值范围**: 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 索引位置。 **取值范围**: 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"ListSqlLimitTaskResponse", string(data)}, " ")
}

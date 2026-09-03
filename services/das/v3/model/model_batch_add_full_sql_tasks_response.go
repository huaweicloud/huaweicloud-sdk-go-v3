package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddFullSqlTasksResponse Response Object
type BatchAddFullSqlTasksResponse struct {

	// 任务ID列表
	Ids *[]int64 `json:"ids,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchAddFullSqlTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddFullSqlTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchAddFullSqlTasksResponse", string(data)}, " ")
}

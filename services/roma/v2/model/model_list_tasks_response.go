package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponse Response Object
type ListTasksResponse struct {

	// 查询任务列表任务总个数
	Total *int64 `json:"total,omitempty"`

	// 查询任务列表返回的当前页的任务个数
	Size *int32 `json:"size,omitempty"`

	// 查询任务列表返回的对象
	Entities       *[]BriefTaskRespBean `json:"entities,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}

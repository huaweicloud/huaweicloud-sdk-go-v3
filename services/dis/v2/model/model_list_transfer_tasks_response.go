package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransferTasksResponse Response Object
type ListTransferTasksResponse struct {

	// 转储任务总数。
	TotalNumber *int32 `json:"total_number,omitempty"`

	// 可创建的转储任务配额。
	Quota *int32 `json:"quota,omitempty"`

	// 转储任务列表。
	Tasks          *[]TransferTask `json:"tasks,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTransferTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransferTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTransferTasksResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBatchOperationTasksResponse Response Object
type ListBatchOperationTasksResponse struct {

	// 批量操作任务列表。
	Tasks *[]ListBatchOperationTasksItem `json:"tasks,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBatchOperationTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchOperationTasksResponse struct{}"
	}

	return strings.Join([]string{"ListBatchOperationTasksResponse", string(data)}, " ")
}

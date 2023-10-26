package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFsTasksResponse Response Object
type ListFsTasksResponse struct {

	// 任务列表
	Tasks          *[]OneFsTaskResp `json:"tasks,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListFsTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFsTasksResponse struct{}"
	}

	return strings.Join([]string{"ListFsTasksResponse", string(data)}, " ")
}

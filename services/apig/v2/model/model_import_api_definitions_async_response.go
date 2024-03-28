package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportApiDefinitionsAsyncResponse Response Object
type ImportApiDefinitionsAsyncResponse struct {

	// 任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportApiDefinitionsAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsAsyncResponse struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsAsyncResponse", string(data)}, " ")
}

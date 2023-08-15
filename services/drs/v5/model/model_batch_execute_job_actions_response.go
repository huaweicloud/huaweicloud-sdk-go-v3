package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExecuteJobActionsResponse Response Object
type BatchExecuteJobActionsResponse struct {

	// 批量异步操作任务响应体。
	Jobs           *[]AsyncActionResp `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchExecuteJobActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteJobActionsResponse struct{}"
	}

	return strings.Join([]string{"BatchExecuteJobActionsResponse", string(data)}, " ")
}

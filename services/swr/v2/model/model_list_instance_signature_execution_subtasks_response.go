package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSignatureExecutionSubtasksResponse Response Object
type ListInstanceSignatureExecutionSubtasksResponse struct {

	// 签名策略执行记录子任务列表
	SubTasks *[]SignatureExecutionSubTask `json:"sub_tasks,omitempty"`

	// 签名策略执行记录子任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceSignatureExecutionSubtasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSignatureExecutionSubtasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSignatureExecutionSubtasksResponse", string(data)}, " ")
}

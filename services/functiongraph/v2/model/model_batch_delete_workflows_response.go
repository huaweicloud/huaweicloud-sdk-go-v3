package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWorkflowsResponse Response Object
type BatchDeleteWorkflowsResponse struct {

	// 成功流程URN列表
	Success *[]string `json:"success,omitempty"`

	// 错误流程详情
	Fail           *[]OperateErrorInfo `json:"fail,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchDeleteWorkflowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkflowsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkflowsResponse", string(data)}, " ")
}

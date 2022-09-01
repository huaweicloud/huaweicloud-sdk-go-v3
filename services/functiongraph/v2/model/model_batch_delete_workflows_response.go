package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteWorkflowsResponse struct {

	// 成功流程URN列表
	Success *[]string `json:"success,omitempty" xml:"success"`

	// 错误流程详情
	Fail           *[]OperateErrorInfo `json:"fail,omitempty" xml:"fail"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchDeleteWorkflowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkflowsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkflowsResponse", string(data)}, " ")
}

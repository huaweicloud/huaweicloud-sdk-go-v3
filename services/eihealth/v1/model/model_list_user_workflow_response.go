package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserWorkflowResponse Response Object
type ListUserWorkflowResponse struct {

	// 所查询类型的流程总数
	Workflows *[]WorkflowListDto `json:"workflows,omitempty"`

	// 当前页的流程列表
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ListUserWorkflowResponse", string(data)}, " ")
}

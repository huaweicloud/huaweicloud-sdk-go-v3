package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNextflowWorkflowResponse Response Object
type ListNextflowWorkflowResponse struct {

	// 当前页的流程列表
	Workflows *[]NextflowWorkflowListDto `json:"workflows,omitempty"`

	// 所查询类型的流程总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNextflowWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNextflowWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ListNextflowWorkflowResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWorkflowResponse struct {

	// 工作流详情集合。
	Elements *[]Workflow `json:"elements,omitempty"`

	// 总数
	TotalElements  *int64 `json:"total_elements,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowResponse", string(data)}, " ")
}

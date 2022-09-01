package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWorkflowExecutionsResponse struct {

	// 流程执行信息列表
	Executions     *[]FlowExecutionBrief `json:"executions,omitempty" xml:"executions"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListWorkflowExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowExecutionsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowExecutionsResponse Response Object
type ListWorkflowExecutionsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 返回个数。
	Count *int32 `json:"count,omitempty"`

	// execution数组。
	Items          *[]WorkflowExecutionResp `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListWorkflowExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowExecutionsResponse", string(data)}, " ")
}

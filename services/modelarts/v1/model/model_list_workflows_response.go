package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowsResponse Response Object
type ListWorkflowsResponse struct {

	// 查询到当前用户名下的所有Workflow总数。
	Total *int32 `json:"total,omitempty"`

	// 查询到当前用户名下的所有符合查询条件的Workflow总数。
	Count *int32 `json:"count,omitempty"`

	// 查询到当前用户名下的所有符合查询条件的Workflow详情。
	Items          *[]Workflow `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListWorkflowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowsResponse", string(data)}, " ")
}

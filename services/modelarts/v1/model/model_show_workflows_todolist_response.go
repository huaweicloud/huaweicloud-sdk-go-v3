package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowsTodolistResponse Response Object
type ShowWorkflowsTodolistResponse struct {

	// 待办列表。
	Items *[]WorkflowTodo `json:"items,omitempty"`

	// 条目个数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWorkflowsTodolistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowsTodolistResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowsTodolistResponse", string(data)}, " ")
}

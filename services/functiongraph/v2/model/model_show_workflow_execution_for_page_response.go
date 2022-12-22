package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWorkflowExecutionForPageResponse struct {

	// 返回所有满足条件的对象个数
	Total *int64 `json:"total,omitempty"`

	// 返回对象的大小
	Size *int32 `json:"size,omitempty"`

	// 函数流返回体信息
	Executions     *[]FlowExecutionBriefV2 `json:"executions,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowWorkflowExecutionForPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowExecutionForPageResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowExecutionForPageResponse", string(data)}, " ")
}

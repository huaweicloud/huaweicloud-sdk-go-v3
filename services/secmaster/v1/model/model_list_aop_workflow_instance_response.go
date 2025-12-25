package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAopWorkflowInstanceResponse Response Object
type ListAopWorkflowInstanceResponse struct {

	// **参数解释**: 流程实例的个数 **约束限制**: 不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**: 流程实例详情的列表 **约束限制**: 不涉及
	Instances      *[]AopworkflowInstanceDetail `json:"instances,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListAopWorkflowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAopWorkflowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListAopWorkflowInstanceResponse", string(data)}, " ")
}

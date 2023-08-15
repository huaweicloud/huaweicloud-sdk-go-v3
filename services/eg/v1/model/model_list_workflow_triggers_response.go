package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowTriggersResponse Response Object
type ListWorkflowTriggersResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]SubscriptionInfo `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListWorkflowTriggersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowTriggersResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowTriggersResponse", string(data)}, " ")
}

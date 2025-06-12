package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackSetOperationsResponse Response Object
type ListStackSetOperationsResponse struct {

	// 资源栈集操作列表
	StackSetOperations *[]StackSetOperation `json:"stack_set_operations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStackSetOperationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackSetOperationsResponse struct{}"
	}

	return strings.Join([]string{"ListStackSetOperationsResponse", string(data)}, " ")
}

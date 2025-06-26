package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSignPolicyExecutionsResponse Response Object
type ListInstanceSignPolicyExecutionsResponse struct {

	// 执行记录列表
	Executions *[]Execution `json:"executions,omitempty"`

	// 执行记录总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceSignPolicyExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSignPolicyExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSignPolicyExecutionsResponse", string(data)}, " ")
}

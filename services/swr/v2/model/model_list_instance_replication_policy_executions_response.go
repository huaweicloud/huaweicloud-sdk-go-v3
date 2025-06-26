package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceReplicationPolicyExecutionsResponse Response Object
type ListInstanceReplicationPolicyExecutionsResponse struct {

	// 执行记录列表
	Executions *[]Execution `json:"executions,omitempty"`

	// 执行记录总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceReplicationPolicyExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceReplicationPolicyExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceReplicationPolicyExecutionsResponse", string(data)}, " ")
}

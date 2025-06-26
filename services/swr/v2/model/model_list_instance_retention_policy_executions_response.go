package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRetentionPolicyExecutionsResponse Response Object
type ListInstanceRetentionPolicyExecutionsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 老化策略执行列表
	Executions     *[]RetentionExecution `json:"executions,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListInstanceRetentionPolicyExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRetentionPolicyExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceRetentionPolicyExecutionsResponse", string(data)}, " ")
}

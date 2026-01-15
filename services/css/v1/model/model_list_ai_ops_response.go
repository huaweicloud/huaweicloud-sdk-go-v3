package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiOpsResponse Response Object
type ListAiOpsResponse struct {
	AiopsList *[]AiOps `json:"aiops_list,omitempty"`

	// 参数解释： 集群风险检测任务总数。 取值范围： 不涉及
	TotalSize      *int32 `json:"total_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAiOpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiOpsResponse struct{}"
	}

	return strings.Join([]string{"ListAiOpsResponse", string(data)}, " ")
}

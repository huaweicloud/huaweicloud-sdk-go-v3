package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceStatisticsRequest Request Object
type ListInstanceStatisticsRequest struct {

	// Agent 类型
	AiAgentType *string `json:"ai_agent_type,omitempty"`
}

func (o ListInstanceStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceStatisticsRequest", string(data)}, " ")
}

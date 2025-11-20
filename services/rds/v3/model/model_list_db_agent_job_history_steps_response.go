package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobHistoryStepsResponse Response Object
type ListDbAgentJobHistoryStepsResponse struct {

	// 执行历史步骤列表。
	Steps *[]ListDbAgentJobHistoryStepsResult `json:"steps,omitempty"`

	// 执行历史步骤总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbAgentJobHistoryStepsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobHistoryStepsResponse struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobHistoryStepsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotJobsResponse Response Object
type ListAutopilotJobsResponse struct {

	// **参数解释**： Job详情列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Body           *[]V2Job `json:"body,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListAutopilotJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotJobsResponse struct{}"
	}

	return strings.Join([]string{"ListAutopilotJobsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobResponse Response Object
type ListJobResponse struct {

	// 作业列表
	Jobs *[]JobListDto `json:"jobs,omitempty"`

	// 作业总数
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 运行中作业总数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RunningCount *int32 `json:"running_count,omitempty"`

	// **参数解释**： 等待中作业总数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	WaitingCount   *int32 `json:"waiting_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResponse struct{}"
	}

	return strings.Join([]string{"ListJobResponse", string(data)}, " ")
}

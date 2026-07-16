package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsResponse Response Object
type ListJobsResponse struct {

	// **参数解释**： 结果总数。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：job列表。
	Items          *[]PoolJob `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}

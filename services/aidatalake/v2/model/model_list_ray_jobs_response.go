package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRayJobsResponse Response Object
type ListRayJobsResponse struct {

	// **参数解释**：下一页查询marker值，若为空表示当前已是最后一页。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**：返回前一页查询地址，为空则表示是第一页。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// **参数解释**：本次查询记录总数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CurrentCount *int32 `json:"current_count,omitempty"`

	// **参数解释**：Ray作业信息列表，包含作业名称、状态、配置等详细信息。数组中的每个元素为一个RayJobInfo对象。 **约束限制**：0~100。
	Jobs           *[]RayJobInfo `json:"jobs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListRayJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRayJobsResponse struct{}"
	}

	return strings.Join([]string{"ListRayJobsResponse", string(data)}, " ")
}

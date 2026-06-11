package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesResponse Response Object
type ListRecycleInstancesResponse struct {

	// **参数解释：** 总记录数。 **取值范围：** 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释：** 实例信息。 **取值范围：** 不涉及。
	Instances      *[]RecycleInstance `json:"instances,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListRecycleInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesResponse", string(data)}, " ")
}

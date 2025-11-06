package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectableResourcesResponse Response Object
type ListProtectableResourcesResponse struct {

	// **参数解释：** 详细的可防护资源的信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items *[]ProtectableResources `json:"items,omitempty"`

	// **参数解释：** 全部可防护资源的数量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProtectableResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListProtectableResourcesResponse", string(data)}, " ")
}

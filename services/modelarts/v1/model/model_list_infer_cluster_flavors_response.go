package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferClusterFlavorsResponse Response Object
type ListInferClusterFlavorsResponse struct {

	// **参数解释：** 当前页码。 **取值范围：** 不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释：** 规格列表。 **取值范围：** 不涉及。
	Data *[]InferFlavor `json:"data,omitempty"`

	// **参数解释：** 总页数。 **取值范围：** 不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释：** 每页数量。 **取值范围：** 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释：** 总记录数。 **取值范围：** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInferClusterFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferClusterFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListInferClusterFlavorsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferClusterFlavorsRequest Request Object
type ListInferClusterFlavorsRequest struct {

	// **参数解释：** 规格类型。 **约束限制：** 不涉及。 **取值范围：** - CPU - GPU - ASCEND **默认取值：** 不涉及。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInferClusterFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferClusterFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListInferClusterFlavorsRequest", string(data)}, " ")
}

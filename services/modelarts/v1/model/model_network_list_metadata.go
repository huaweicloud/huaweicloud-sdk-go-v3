package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkListMetadata struct {

	// **参数解释**：分页查询时，下一次查询位置。 **取值范围**：不涉及。
	Continue *string `json:"continue,omitempty"`

	// **参数解释**：剩余资源个数。 **取值范围**：不涉及。
	RemainingItemCount *int32 `json:"remainingItemCount,omitempty"`
}

func (o NetworkListMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkListMetadata struct{}"
	}

	return strings.Join([]string{"NetworkListMetadata", string(data)}, " ")
}

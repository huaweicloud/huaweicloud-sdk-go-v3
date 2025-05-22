package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusStatistics **参数解释**： 资源统计信息。 **取值范围**： 不涉及。
type StatusStatistics struct {

	// **参数解释**： 活跃资源。 **取值范围**： 不涉及。
	Active *int64 `json:"active,omitempty"`

	// **参数解释**： 总资源。 **取值范围**： 不涉及。
	Total *int64 `json:"total,omitempty"`
}

func (o StatusStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusStatistics struct{}"
	}

	return strings.Join([]string{"StatusStatistics", string(data)}, " ")
}

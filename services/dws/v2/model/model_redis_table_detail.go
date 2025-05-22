package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisTableDetail **参数解释**： 重分布表详情。 **取值范围**： 不涉及。
type RedisTableDetail struct {

	// **参数解释**： 重分布具体数据。 **取值范围**： 不涉及。
	Data *[]RedisTable `json:"data,omitempty"`

	// **参数解释**： 重分布表张数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o RedisTableDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisTableDetail struct{}"
	}

	return strings.Join([]string{"RedisTableDetail", string(data)}, " ")
}

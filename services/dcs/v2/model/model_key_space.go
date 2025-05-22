package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeySpace 键值空间
type KeySpace struct {

	// **参数解释**： db索引值。 **取值范围**： 不涉及。
	DbIdx *string `json:"db_idx,omitempty"`

	// **参数解释**： 节点键数量。 **取值范围**： 不涉及。
	Keys *string `json:"keys,omitempty"`

	// **参数解释**： 节点过期键数量。 **取值范围**： 不涉及。
	Expires *string `json:"expires,omitempty"`

	// **参数解释**： 节点键的平均过期时间。 **取值范围**： 不涉及。
	AvgTtl *string `json:"avg_ttl,omitempty"`
}

func (o KeySpace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeySpace struct{}"
	}

	return strings.Join([]string{"KeySpace", string(data)}, " ")
}

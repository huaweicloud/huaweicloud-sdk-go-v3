package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModelNetwork **参数解释**：资源池网络参数。
type PoolSpecModelNetwork struct {

	// **参数解释**：网络ID。 **取值范围**：不涉及。
	Name string `json:"name"`
}

func (o PoolSpecModelNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelNetwork struct{}"
	}

	return strings.Join([]string{"PoolSpecModelNetwork", string(data)}, " ")
}

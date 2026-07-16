package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Pool struct {

	// **参数解释**：专属资源池ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：专属资源池名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o Pool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}

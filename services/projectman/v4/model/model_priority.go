package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Priority 工作项优先级
type Priority struct {

	// **参数解释：** 工作项的优先级。 **取值范围：** - 低。 - 中。 - 高。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 优先级id 。 **取值范围：** 1（低） 2（中） 3（高）
	Id *int32 `json:"id,omitempty"`
}

func (o Priority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Priority struct{}"
	}

	return strings.Join([]string{"Priority", string(data)}, " ")
}

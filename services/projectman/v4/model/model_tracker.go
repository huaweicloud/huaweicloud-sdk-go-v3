package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tracker 工作项类型
type Tracker struct {

	// **参数解释：** 类型名称。 **取值范围：** - Task。 - Bug。 - Epic。 - Feature。 - Story。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 类型id。 **取值范围：** 2（任务/Task） 3（缺陷/Bug） 5（Epic） 6（Feature） 7（Story）
	Id *int32 `json:"id,omitempty"`
}

func (o Tracker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tracker struct{}"
	}

	return strings.Join([]string{"Tracker", string(data)}, " ")
}

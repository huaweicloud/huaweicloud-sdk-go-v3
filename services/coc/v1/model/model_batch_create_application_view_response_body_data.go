package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateApplicationViewResponseBodyData **参数解释：** 创建成功的的应用、组件和分组id列表组成列表 **取值范围：** 大小在0到500之间。
type BatchCreateApplicationViewResponseBodyData struct {

	// **参数解释：** 应用id列表。 **取值范围：** 不涉及。
	ApplicationIds *[]string `json:"application_ids,omitempty"`

	// **参数解释：** 组件id列表。 **取值范围：** 不涉及。
	ComponentIds *[]string `json:"component_ids,omitempty"`

	// **参数解释：** 分组id列表。 **取值范围：** 不涉及。
	GroupIds *[]string `json:"group_ids,omitempty"`
}

func (o BatchCreateApplicationViewResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewResponseBodyData struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewResponseBodyData", string(data)}, " ")
}

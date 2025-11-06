package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessProgress struct {

	// **参数解释：** 接入步骤（1:回源IP加白步骤，2:本地验证步骤 3:修改DNS解析步骤） **约束限制：** 不涉及 **取值范围：** - 1 : 回源IP加白步骤 - 2 : 本地验证步骤 - 3 : 修改DNS解析步骤  **默认取值：** 不涉及
	Step *int32 `json:"step,omitempty"`

	// **参数解释：** 完成情况（0:未完成，1:完成，2:跳过） **约束限制：** 不涉及 **取值范围：** - 0:未完成 - 1:完成 - 2:跳过  **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`
}

func (o AccessProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessProgress struct{}"
	}

	return strings.Join([]string{"AccessProgress", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SkippedCheckItemList struct {

	// **参数解释：** 跳过检查的项目名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	ResourceSelector *ResourceSelector `json:"resourceSelector,omitempty"`
}

func (o SkippedCheckItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkippedCheckItemList struct{}"
	}

	return strings.Join([]string{"SkippedCheckItemList", string(data)}, " ")
}

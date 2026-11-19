package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyOpsWindowRequestBody struct {

	// **参数解释**: 可维护时间段开始时间。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 可维护时间段结束时间。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTime string `json:"end_time"`
}

func (o ModifyOpsWindowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyOpsWindowRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyOpsWindowRequestBody", string(data)}, " ")
}

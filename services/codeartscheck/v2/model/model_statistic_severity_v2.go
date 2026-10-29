package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticSeverityV2 struct {

	// **参数解释：** 致命问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Critical *int32 `json:"critical,omitempty"`

	// **参数解释：** 严重问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Major *int32 `json:"major,omitempty"`

	// **参数解释：** 一般问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Minor *int32 `json:"minor,omitempty"`

	// **参数解释：** 提示问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Suggestion *int32 `json:"suggestion,omitempty"`
}

func (o StatisticSeverityV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticSeverityV2 struct{}"
	}

	return strings.Join([]string{"StatisticSeverityV2", string(data)}, " ")
}

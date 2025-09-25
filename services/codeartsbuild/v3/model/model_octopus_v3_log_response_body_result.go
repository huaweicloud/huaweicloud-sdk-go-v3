package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OctopusV3LogResponseBodyResult 结果
type OctopusV3LogResponseBodyResult struct {

	// **参数解释**： 是否还有剩余日志标识。 **约束限制**： 不涉及。 **取值范围**： true或false。 **默认取值**： 不涉及。
	HasMore *bool `json:"has_more,omitempty"`

	// **参数解释**： 日志查询起始偏移量。 **约束限制**： 不涉及。 **取值范围**： 数字组成。 **默认取值**： 不涉及。
	StartOffset *string `json:"start_offset,omitempty"`

	// **参数解释**： 日志查询结束偏移量。 **约束限制**： 不涉及。 **取值范围**： 数字组成。 **默认取值**： 不涉及。
	EndOffset *string `json:"end_offset,omitempty"`

	// **参数解释**： 返回日志内容，可能会空，请再次请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Log *string `json:"log,omitempty"`

	// **参数解释**： 日志来源。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Location *string `json:"location,omitempty"`
}

func (o OctopusV3LogResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OctopusV3LogResponseBodyResult struct{}"
	}

	return strings.Join([]string{"OctopusV3LogResponseBodyResult", string(data)}, " ")
}

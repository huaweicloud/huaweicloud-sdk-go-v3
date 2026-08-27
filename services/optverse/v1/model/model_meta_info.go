package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaInfo **参数解释**： 补充元信息。 **约束限制**： 不涉及 **取值范围**： 不涉及。 **默认取值**： 不涉及
type MetaInfo struct {

	// **参数解释**： 请求用时。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	CostTime *int64 `json:"cost_time,omitempty"`

	// **参数解释**： 请求结束时间。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	CurrentTime *int64 `json:"current_time,omitempty"`
}

func (o MetaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaInfo struct{}"
	}

	return strings.Join([]string{"MetaInfo", string(data)}, " ")
}

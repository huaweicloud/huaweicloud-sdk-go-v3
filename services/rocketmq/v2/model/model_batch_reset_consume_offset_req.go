package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResetConsumeOffsetReq struct {

	// **参数解释**： 消费组列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Groups *[]string `json:"groups,omitempty"`

	// **参数解释**： 重置的时间。-1表示重置到最新位点。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o BatchResetConsumeOffsetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetConsumeOffsetReq struct{}"
	}

	return strings.Join([]string{"BatchResetConsumeOffsetReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackTrendRespBody struct {

	// **参数解释**： 阻断次数 **取值范围**： 不涉及
	DenyCount *int64 `json:"deny_count,omitempty"`

	// **参数解释**： 放行次数 **取值范围**： 不涉及
	PermitCount *int64 `json:"permit_count,omitempty"`

	// **参数解释**： 聚合时间 **取值范围**： 不涉及
	Time *int64 `json:"time,omitempty"`
}

func (o AttackTrendRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackTrendRespBody struct{}"
	}

	return strings.Join([]string{"AttackTrendRespBody", string(data)}, " ")
}

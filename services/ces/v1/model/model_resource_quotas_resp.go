package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceQuotasResp **参数解释**： 资源配额。
type ResourceQuotasResp struct {

	// **参数解释**： 配额类型。 **取值范围**： 枚举值说明：   - alarm：告警规则
	Type string `json:"type"`

	// **参数解释**： 已使用配额数。 **取值范围**： 不涉及。
	Used int64 `json:"used"`

	// **参数解释**： 单位。 **取值范围**： 长度为[0,32]个字符。
	Unit string `json:"unit"`

	// **参数解释**： 配额总数。 **取值范围**： 不涉及。
	Quota int64 `json:"quota"`

	// **参数解释**： 最小值。 **取值范围**： 不涉及。
	Min *int64 `json:"min,omitempty"`

	// **参数解释**： 最大值。 **取值范围**： 不涉及。
	Max *int64 `json:"max,omitempty"`
}

func (o ResourceQuotasResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceQuotasResp struct{}"
	}

	return strings.Join([]string{"ResourceQuotasResp", string(data)}, " ")
}

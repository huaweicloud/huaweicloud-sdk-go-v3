package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaselineInfo struct {

	// **参数解释**： 基线ID。 **取值范围**： 长度为1~64个字符。
	BaselineId *string `json:"baseline_id,omitempty"`

	// **参数解释**： 基线名称。 **取值范围**： 长度为1~1024个字符。
	BaselineName *string `json:"baseline_name,omitempty"`

	// **参数解释**： 基线承诺时间，单位毫秒。 **取值范围**： 不涉及。
	SlaTime *int64 `json:"sla_time,omitempty"`

	// **参数解释**： 基线预警时间，单位毫秒。 **取值范围**： 不涉及。
	ExpTime *int64 `json:"exp_time,omitempty"`
}

func (o BaselineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineInfo struct{}"
	}

	return strings.Join([]string{"BaselineInfo", string(data)}, " ")
}

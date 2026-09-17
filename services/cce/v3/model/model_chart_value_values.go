package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChartValueValues **参数解释：** values.yaml中的数据，数据结构以具体的模板为准。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type ChartValueValues struct {
	Basic *interface{} `json:"basic,omitempty"`
}

func (o ChartValueValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChartValueValues struct{}"
	}

	return strings.Join([]string{"ChartValueValues", string(data)}, " ")
}

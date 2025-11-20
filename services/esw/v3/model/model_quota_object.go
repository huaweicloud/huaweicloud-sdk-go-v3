package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaObject 配额信息
type QuotaObject struct {

	// - 参数解释：配额类型。 - 约束限制：不涉及。 - 取值范围：instance。 - 默认取值：不涉及。
	Type string `json:"type"`

	// - 参数解释：配额总量。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Quota int32 `json:"quota"`

	// - 参数解释：已使用配额。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Used int32 `json:"used"`

	// - 参数解释：配额总量的最小值。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Min int32 `json:"min"`

	// - 参数解释：配额总量的最大值。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Max int32 `json:"max"`
}

func (o QuotaObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaObject struct{}"
	}

	return strings.Join([]string{"QuotaObject", string(data)}, " ")
}

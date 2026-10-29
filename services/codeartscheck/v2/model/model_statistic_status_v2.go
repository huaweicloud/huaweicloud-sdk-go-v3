package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticStatusV2 struct {

	// **参数解释：** [待处理问题数。](tag:hws,hws_hk)[待处理和重新打开（需打开忽略审核开关）问题数。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm) **取值范围：** 不涉及。
	Unresolved *int32 `json:"unresolved,omitempty"`

	// **参数解释：** 已解决和已修复问题数。 **取值范围：** 不涉及。
	Resolved *int32 `json:"resolved,omitempty"`

	// **参数解释：** 已忽略问题数。 **取值范围：** 不涉及。
	Dismissed *int32 `json:"dismissed,omitempty"`

	// **参数解释：** 已修复问题数。 **取值范围：** 不涉及。
	Fixed *int32 `json:"fixed,omitempty"`

	// **参数解释：** 已解决问题数。 **取值范围：** 不涉及。
	Solved *int32 `json:"solved,omitempty"`
}

func (o StatisticStatusV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticStatusV2 struct{}"
	}

	return strings.Join([]string{"StatisticStatusV2", string(data)}, " ")
}

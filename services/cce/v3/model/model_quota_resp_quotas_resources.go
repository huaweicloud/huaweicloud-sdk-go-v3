package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaRespQuotasResources struct {

	// **参数解释：** 类型 **约束限制：** 不涉及 **取值范围：** - Charts：配额类型为模板  **默认取值：** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释：** 配额 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释：** 已使用量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaRespQuotasResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaRespQuotasResources struct{}"
	}

	return strings.Join([]string{"QuotaRespQuotasResources", string(data)}, " ")
}

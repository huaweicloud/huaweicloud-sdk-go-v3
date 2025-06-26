package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasResource **参数解释**： 资源配额详情。 **取值范围**： 不涉及。
type QuotasResource struct {

	// **参数解释**： 项目资源类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 已使用的资源数量。 **取值范围**： 不涉及。
	Used int32 `json:"used"`

	// **参数解释**： 项目资源配额。 **取值范围**： 不涉及。
	Quota int32 `json:"quota"`

	// **参数解释**： 资源计量单位。 **取值范围**： 不涉及。
	Unit string `json:"unit"`
}

func (o QuotasResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasResource struct{}"
	}

	return strings.Join([]string{"QuotasResource", string(data)}, " ")
}

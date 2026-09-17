package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaResource struct {

	// **参数解释：** 资源类型 **约束限制：** 不涉及 **取值范围：** - cluster：表示集群配额  **默认取值：** 不涉及
	QuotaKey *string `json:"quotaKey,omitempty"`

	// **参数解释：** 资源配额值 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	QuotaLimit *int32 `json:"quotaLimit,omitempty"`

	// **参数解释：** 已创建的资源个数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Used *int32 `json:"used,omitempty"`

	// **参数解释：** 资源单位 **约束限制：** 不涉及 **取值范围：** - count：个数  **默认取值：** 不涉及
	Unit *string `json:"unit,omitempty"`

	// **参数解释：** 局点ID。若资源不涉及此参数，则不返回该参数。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RegionId *string `json:"regionId,omitempty"`

	// **参数解释：** 可用区ID。若资源不涉及此参数，则不返回该参数。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AvailabilityZoneId *string `json:"availabilityZoneId,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}

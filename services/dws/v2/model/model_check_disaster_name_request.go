package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDisasterNameRequest Request Object
type CheckDisasterNameRequest struct {

	// **参数解释**： 容灾名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DrName string `json:"dr_name"`

	// **参数解释**： 容灾类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 备集群所在局点。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StandbyRegion *string `json:"standby_region,omitempty"`

	// **参数解释**： 备集群所在项目ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StandbyProjectId *string `json:"standby_project_id,omitempty"`
}

func (o CheckDisasterNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDisasterNameRequest struct{}"
	}

	return strings.Join([]string{"CheckDisasterNameRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OccupyResource **参数解释**： 用户占用资源列表。 **取值范围**： 不涉及。
type OccupyResource struct {

	// **参数解释**： 资源项名称。 **取值范围**： 不涉及。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 资源属性数值。 **取值范围**： 不涉及。
	ResourceValue *int32 `json:"resource_value,omitempty"`

	// **参数解释**： 资源属性单位。 **取值范围**： 不涉及。
	ValueUnit *string `json:"value_unit,omitempty"`

	// **参数解释**： 资源附加描述。 **取值范围**： 不涉及。
	ResourceDescription *string `json:"resource_description,omitempty"`
}

func (o OccupyResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OccupyResource struct{}"
	}

	return strings.Join([]string{"OccupyResource", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenPublicIp **参数解释**： 弹性公网IP对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type OpenPublicIp struct {

	// **参数解释**： 弹性IP绑定类型。 **约束限制**： 不涉及。 **取值范围**： - auto_assign：自动绑定 - not_use：暂未使用 - bind_existing ：使用已有  **默认取值**： 不涉及。
	PublicBindType *string `json:"public_bind_type,omitempty"`

	// **参数解释**： 弹性IP的ID **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EipId *string `json:"eip_id,omitempty"`
}

func (o OpenPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenPublicIp struct{}"
	}

	return strings.Join([]string{"OpenPublicIp", string(data)}, " ")
}

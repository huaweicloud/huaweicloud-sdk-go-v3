package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartInstanceResponse Response Object
type RestartInstanceResponse struct {

	// **参数解释**： 实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： node列表 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Nodes *[]string `json:"nodes,omitempty"`

	// **参数解释**： 结果。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartInstanceResponse", string(data)}, " ")
}

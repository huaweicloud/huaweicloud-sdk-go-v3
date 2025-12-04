package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupCreateReq struct {

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 消费组描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及
	GroupDesc *string `json:"group_desc,omitempty"`
}

func (o GroupCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupCreateReq struct{}"
	}

	return strings.Join([]string{"GroupCreateReq", string(data)}, " ")
}

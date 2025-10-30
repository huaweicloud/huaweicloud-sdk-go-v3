package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoNodeExpansionPolicyRequest Request Object
type ShowAutoNodeExpansionPolicyRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 语言。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAutoNodeExpansionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoNodeExpansionPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoNodeExpansionPolicyRequest", string(data)}, " ")
}

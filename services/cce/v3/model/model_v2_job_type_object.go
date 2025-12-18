package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type V2JobTypeObject struct {

	// **参数解释**： Job UUID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： Job 创建时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**： Job 最后更新时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o V2JobTypeObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2JobTypeObject struct{}"
	}

	return strings.Join([]string{"V2JobTypeObject", string(data)}, " ")
}

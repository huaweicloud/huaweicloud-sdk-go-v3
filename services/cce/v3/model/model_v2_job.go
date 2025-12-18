package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type V2Job struct {

	// **参数解释**： API类型 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： Job
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *V2JobTypeObject `json:"metadata,omitempty"`

	Spec *V2JobSpec `json:"spec,omitempty"`

	Status *V2JobStatus `json:"status,omitempty"`
}

func (o V2Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2Job struct{}"
	}

	return strings.Join([]string{"V2Job", string(data)}, " ")
}

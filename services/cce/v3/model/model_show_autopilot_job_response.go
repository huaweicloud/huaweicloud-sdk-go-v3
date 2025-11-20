package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotJobResponse Response Object
type ShowAutopilotJobResponse struct {

	// **参数解释**： API类型 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： Job
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *JobMetadata `json:"metadata,omitempty"`

	Spec *JobSpec `json:"spec,omitempty"`

	Status         *JobStatus `json:"status,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowAutopilotJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotJobResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotJobResponse", string(data)}, " ")
}

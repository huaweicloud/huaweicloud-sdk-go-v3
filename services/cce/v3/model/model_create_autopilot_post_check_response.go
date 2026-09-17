package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutopilotPostCheckResponse Response Object
type CreateAutopilotPostCheckResponse struct {

	// **参数解释：** API版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 资源类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	Metadata *PostcheckCluserResponseMetadata `json:"metadata,omitempty"`

	Spec *PostcheckSpec `json:"spec,omitempty"`

	Status         *PostcheckClusterResponseBodyStatus `json:"status,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o CreateAutopilotPostCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutopilotPostCheckResponse struct{}"
	}

	return strings.Join([]string{"CreateAutopilotPostCheckResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareConfigurationResponse Response Object
type CompareConfigurationResponse struct {

	// **参数解释**：  源参数模板id。  **约束限制**：  不涉及。
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**：  目标参数模板id。  **约束限制**：  不涉及。
	TargetId *string `json:"target_id,omitempty"`

	// **参数解释**：  源参数模板名称。  **约束限制**：  不涉及。
	SourceName *string `json:"source_name,omitempty"`

	// **参数解释**：  目标参数模板名称。  **约束限制**：  不涉及。
	TargetName *string `json:"target_name,omitempty"`

	// **参数解释**：  模板参数差异信息。  **约束限制**：  不涉及。
	Parameters     *[]ParamGroupParameterDiff `json:"parameters,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CompareConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CompareConfigurationResponse", string(data)}, " ")
}

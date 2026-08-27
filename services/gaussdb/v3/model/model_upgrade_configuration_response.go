package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeConfigurationResponse Response Object
type UpgradeConfigurationResponse struct {

	// **参数解释**：  参数模板名称。  **取值范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  差异参数列表。
	DiffParameters *[]GroupParameterDiffInfo `json:"diff_parameters,omitempty"`

	// **参数解释**：  执行更新操作被跳过的参数名称列表（原值与目标值相同）。
	SkippedParameterNames *[]string `json:"skipped_parameter_names,omitempty"`
	HttpStatusCode        int       `json:"-"`
}

func (o UpgradeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpgradeConfigurationResponse", string(data)}, " ")
}

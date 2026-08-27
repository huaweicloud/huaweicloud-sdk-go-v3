package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeConfigurationRequestBody **参数解释**：  更新自定义模板请求体。
type UpgradeConfigurationRequestBody struct {

	// **参数解释**：  需要更新的差异参数名称列表。 - 若参数有值传入：将该参数更新为系统默认模板的值。 - 若参数传入空值或未传入：保留自定义模板中的原有值。  **约束限制**：  不涉及。
	Parameters *[]string `json:"parameters,omitempty"`
}

func (o UpgradeConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeConfigurationRequestBody", string(data)}, " ")
}

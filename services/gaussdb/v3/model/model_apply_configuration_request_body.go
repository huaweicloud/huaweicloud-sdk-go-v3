package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyConfigurationRequestBody struct {

	// 实例ID列表。列表长度限制在10以内。
	InstanceIds []string `json:"instance_ids"`

	// **参数解释**：  是否更新实例参数组版本，更新后实例规格变更时默认的规格参数值会以最新版本的为准。  **约束限制**：  不涉及。  **取值范围**：  - true：是。 - false：否。  **默认取值**：    false。
	IsUpdateParamGroupVersion *bool `json:"is_update_param_group_version,omitempty"`
}

func (o ApplyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequestBody", string(data)}, " ")
}

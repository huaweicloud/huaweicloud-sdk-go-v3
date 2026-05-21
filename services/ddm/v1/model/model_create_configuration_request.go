package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationRequest struct {

	// 名称。
	Name string `json:"name"`

	// 描述。
	Description string `json:"description"`

	// **参数解释**：  参数值对象，用户基于默认参数模板自定义的参数值。  **约束限制**：  不涉及。  **取值范围**：  - key：参数名称，如“contains_shard_key”，“connection_idle_timeout”。为空时不修改参数值。  - value：参数值，如“6”，“20”。key不为空时value也不可为空。  **默认取值**：  不涉及。
	Values map[string]string `json:"values,omitempty"`
}

func (o CreateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequest", string(data)}, " ")
}

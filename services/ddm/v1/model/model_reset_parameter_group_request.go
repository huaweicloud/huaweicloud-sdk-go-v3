package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetParameterGroupRequest Request Object
type ResetParameterGroupRequest struct {

	// **参数解释**：  参数组ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	ConfigId string `json:"config_id"`

	Body *ConfigurationUpdateReqV3 `json:"body,omitempty"`
}

func (o ResetParameterGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetParameterGroupRequest struct{}"
	}

	return strings.Join([]string{"ResetParameterGroupRequest", string(data)}, " ")
}

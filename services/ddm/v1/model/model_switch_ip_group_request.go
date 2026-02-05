package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchIpGroupRequest Request Object
type SwitchIpGroupRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  组ID，此参数是组的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成。  **默认取值**：  不涉及。
	GroupId string `json:"group_id"`

	Body *ElbIpGroupOpsReq `json:"body,omitempty"`
}

func (o SwitchIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchIpGroupRequest struct{}"
	}

	return strings.Join([]string{"SwitchIpGroupRequest", string(data)}, " ")
}

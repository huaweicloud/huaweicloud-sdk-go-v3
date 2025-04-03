package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyNodeDataIpRequest struct {

	// **参数解释**：              内网IP地址。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认值**：  不涉及。
	InternalIp string `json:"internal_ip"`
}

func (o ModifyNodeDataIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyNodeDataIpRequest struct{}"
	}

	return strings.Join([]string{"ModifyNodeDataIpRequest", string(data)}, " ")
}

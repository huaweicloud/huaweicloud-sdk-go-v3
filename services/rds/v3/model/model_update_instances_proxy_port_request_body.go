package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstancesProxyPortRequestBody struct {

	// **参数解释**：  端口号。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Port int32 `json:"port"`
}

func (o UpdateInstancesProxyPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstancesProxyPortRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstancesProxyPortRequestBody", string(data)}, " ")
}

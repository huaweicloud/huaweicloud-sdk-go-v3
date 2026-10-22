package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkOperatorStateRequest Request Object
type ShowSparkOperatorStateRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：异步操作唯一ID，用于标识一次异步操作（如重启集群等）。 **约束限制**：不涉及。 **取值范围**：只能由英文字母、数字及中划线组成，长度为1~64个字符。 **默认取值**：不涉及。
	OperationId string `json:"operation_id"`
}

func (o ShowSparkOperatorStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkOperatorStateRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkOperatorStateRequest", string(data)}, " ")
}

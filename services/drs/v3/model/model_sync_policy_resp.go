package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncPolicyResp struct {

	// **参数解释**：  设置同步策略的任务ID，用于标识策略设置结果对应的任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  策略设置结果状态，表示同步策略是否设置成功。  **约束限制**：  不涉及。  **取值范围**：  - success：设置成功 - failed：设置失败  **默认取值**：  不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o SyncPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncPolicyResp struct{}"
	}

	return strings.Join([]string{"SyncPolicyResp", string(data)}, " ")
}

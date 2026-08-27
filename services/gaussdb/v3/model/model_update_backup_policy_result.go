package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBackupPolicyResult struct {

	// **参数解释**：  实例ID。  **取值范围**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  提交备份策略异常时返回的编码。  **取值范围**：  不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  提交备份策略异常时返回的描述信息。  **取值范围**：  不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o UpdateBackupPolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyResult struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyResult", string(data)}, " ")
}

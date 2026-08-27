package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaurusDbAdvancedBackupPolicyResponse Response Object
type UpdateTaurusDbAdvancedBackupPolicyResponse struct {

	// **参数解释**：  状态信息。  **取值范围**：  COMPLETED：设置备份策略成功。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 实例ID，严格匹配UUID规则。 **取值范围**： 与请求的实例ID相同。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 实例ID对应的实例名称。
	InstanceName   *string `json:"instance_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaurusDbAdvancedBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaurusDbAdvancedBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaurusDbAdvancedBackupPolicyResponse", string(data)}, " ")
}

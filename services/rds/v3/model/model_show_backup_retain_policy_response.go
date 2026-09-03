package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRetainPolicyResponse Response Object
type ShowBackupRetainPolicyResponse struct {

	// **参数解释**：  实例id  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  实例名字  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  引擎类型  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	EngineName *string `json:"engine_name,omitempty"`

	// **参数解释**：  实例引擎版本  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**：  实例删除时间  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	InstanceDeleteTime *int64 `json:"instance_delete_time,omitempty"`

	// **参数解释**  自动备份保留策略。NONE不保留，LAST保留最后一个，ALL全部保留。  **约束限制**  不涉及  **取值范围**  NONE、LAST、ALL  **默认取值**  不涉及
	Auto *string `json:"auto,omitempty"`

	// **参数解释**  手动备份保留策略。NONE不保留，LAST保留最后一个，ALL全部保留。  **约束限制**  不涉及  **取值范围**  NONE、LAST、ALL  **默认取值**  不涉及
	Manual         *string `json:"manual,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackupRetainPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRetainPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupRetainPolicyResponse", string(data)}, " ")
}

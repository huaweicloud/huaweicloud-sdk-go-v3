package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoCesAlarmInfo struct {

	// **参数解释**：  告警记录唯一标识符。  **约束限制**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  租户ID。  **约束限制**：  不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**：  租户名称。  **约束限制**：  不涉及。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**：  用户ID。  **约束限制**：  不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：  用户名称。  **约束限制**：  不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：  项目ID。  **约束限制**：  不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  项目名称。  **约束限制**：  不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释**：  数据库引擎名称。  **约束限制**：  不涉及。
	EngineName *string `json:"engine_name,omitempty"`

	// **参数解释**：  新实例是否默认开启自动告警。  **约束限制**：  不涉及。
	NewInstanceDefault *bool `json:"new_instance_default,omitempty"`

	// **参数解释**：  自动告警状态转换。  **约束限制**：  不涉及。
	SwitchStatus *string `json:"switch_status,omitempty"`

	// **参数解释**：  告警规则唯一标识符。  **约束限制**：  不涉及。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**：  主题URN。  **约束限制**：  不涉及。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**：  记录被创建的时间戳。  **约束限制**：  不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**：  记录最后一次被更新的时间戳。  **约束限制**：  不涉及。
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o AutoCesAlarmInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoCesAlarmInfo struct{}"
	}

	return strings.Join([]string{"AutoCesAlarmInfo", string(data)}, " ")
}

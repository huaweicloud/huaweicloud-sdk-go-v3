package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventSubUpdateRequest **参数解释**： 更新订阅事件请求体。 **取值范围**： 不涉及。
type EventSubUpdateRequest struct {

	// **参数解释**： 事件订阅名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 事件源类型。 **取值范围**： - cluster：集群。 - backup：快照。 - disaster-recovery：容灾。 - data.migration：数据迁移。 - dws.ingestion：DwsIngestion。
	SourceType *string `json:"source_type,omitempty"`

	// **参数解释**： 事件源ID。 **取值范围**： 不涉及。
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**： 事件类别。 **取值范围**： 支持management、monitor、security、system alarm。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 事件级别。 **取值范围**： 支持normal、warning。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 事件标签。 **取值范围**： 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 是否开启订阅。 **取值范围**： 1为开启，0为关闭。
	Enable *int32 `json:"enable,omitempty"`

	// **参数解释**： 消息通知地址。 **取值范围**： 不涉及。
	NotificationTarget string `json:"notification_target"`

	// **参数解释**： 消息主题名称。 **取值范围**： 不涉及。
	NotificationTargetName string `json:"notification_target_name"`

	// **参数解释**： 消息主题类型。 **取值范围**： - SMN：SMN类型
	NotificationTargetType string `json:"notification_target_type"`
}

func (o EventSubUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSubUpdateRequest struct{}"
	}

	return strings.Join([]string{"EventSubUpdateRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventResponse **参数解释**： 事件返回体。 **取值范围**： 不涉及。
type EventResponse struct {

	// **参数解释**： 事件类别。 **取值范围**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 事件描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 事件ID。 **取值范围**： 不涉及。
	EventId *string `json:"event_id,omitempty"`

	// **参数解释**： 事件定义名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 事件显示名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 所属服务。 **取值范围**： 不涉及。
	NameSpace *string `json:"name_space,omitempty"`

	// **参数解释**： 事件级别。 **取值范围**： 不涉及。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 事件源类别。 **取值范围**： - cluster：集群。 - backup：快照。 - disaster-recovery：容灾。 - data.migration：数据迁移。 - dws.ingestion：DwsIngestion。
	SourceType *string `json:"source_type,omitempty"`

	// **参数解释**： 时间。 **取值范围**： 不涉及。
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 租户凭证ID。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 事件源ID。 **取值范围**： 不涉及。
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**： 事件源名称。 **取值范围**： 不涉及。
	SourceName *string `json:"source_name,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - 1：已保存。 - 2：已处理。
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 事件主题。 **取值范围**： 不涉及。
	Subject *string `json:"subject,omitempty"`

	// **参数解释**： 事件信息。 **取值范围**： 不涉及。
	Context *string `json:"context,omitempty"`
}

func (o EventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResponse struct{}"
	}

	return strings.Join([]string{"EventResponse", string(data)}, " ")
}

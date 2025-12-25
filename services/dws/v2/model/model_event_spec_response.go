package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventSpecResponse **参数解释**： 事件配置对象。 **取值范围**： 不涉及。
type EventSpecResponse struct {

	// **参数解释**： 事件配置ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 事件配置定义名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 事件配置显示名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 事件配置描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 事件主题。 **取值范围**： 不涉及。
	Subject *string `json:"subject,omitempty"`

	// **参数解释**： 事件类别。 **取值范围**： - management：管理。 - monitor：监控。 - security：安全。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 事件级别。 **取值范围**： 不涉及。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 事件源类型。 **取值范围**： - cluster：集群。 - backup：快照。 - disaster-recovery：容灾。 - data.migration：数据迁移。 - dws.ingestion：DwsIngestion。
	SourceType *string `json:"source_type,omitempty"`

	// **参数解释**： 所属服务。 **取值范围**： 不涉及。
	NameSpace *string `json:"name_space,omitempty"`
}

func (o EventSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSpecResponse struct{}"
	}

	return strings.Join([]string{"EventSpecResponse", string(data)}, " ")
}

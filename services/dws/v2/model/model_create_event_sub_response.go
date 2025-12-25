package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventSubResponse Response Object
type CreateEventSubResponse struct {

	// **参数解释**： 订阅ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 订阅名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 事件源类型。 **取值范围**： - cluster：集群。 - backup：快照。 - disaster-recovery：容灾。 - data.migration：数据迁移。 - dws.ingestion：DwsIngestion。
	SourceType *string `json:"source_type,omitempty"`

	// **参数解释**： 事件源ID。 **取值范围**： 不涉及。
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**： 事件类别。 **取值范围**： - management：管理。 - monitor：监控。 - security：安全。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 事件级别。 **取值范围**： 不涉及。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 事件标签。 **取值范围**： 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 是否开启订阅。 **取值范围**： 1为开启，0为关闭。
	Enable *int32 `json:"enable,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 所属服务。 **取值范围**： 不涉及。
	NameSpace *string `json:"name_space,omitempty"`

	// **参数解释**： 消息通知主题地址。 **取值范围**： 不涉及。
	NotificationTarget *string `json:"notification_target,omitempty"`

	// **参数解释**： 消息通知主题名称。 **取值范围**： 不涉及。
	NotificationTargetName *string `json:"notification_target_name,omitempty"`

	// **参数解释**： 消息主题类型。 **取值范围**： - SMN：SMN类型
	NotificationTargetType *string `json:"notification_target_type,omitempty"`

	// **参数解释**： 语言。 **取值范围**： 不涉及。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 时区。 **取值范围**： 不涉及。
	TimeZone       *string `json:"time_zone,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventSubResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSubResponse struct{}"
	}

	return strings.Join([]string{"CreateEventSubResponse", string(data)}, " ")
}

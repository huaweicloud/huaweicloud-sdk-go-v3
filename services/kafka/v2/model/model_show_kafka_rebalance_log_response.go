package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaRebalanceLogResponse Response Object
type ShowKafkaRebalanceLogResponse struct {

	// **参数解释**： 日志ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 重平衡日志状态。 **取值范围**： - OPEN：已开启重平衡日志。 - CLOSE：已关闭重平衡日志。 - OPENING：重平衡日志开启中。 - CLOSING：重平衡日志关闭中。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 日志流ID。 **取值范围**： 不涉及。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// **参数解释**： 日志组ID。 **取值范围**： 不涉及。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释**： 看板ID。 **取值范围**： 不涉及。
	DashboardId *string `json:"dashboard_id,omitempty"`

	// **参数解释**： 日志类型。 **取值范围**： 不涉及。
	LogType *string `json:"log_type,omitempty"`

	// **参数解释**： 日志文件名称。 **取值范围**： 不涉及。
	LogFileName *string `json:"log_file_name,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	UpdatedAt      *int64 `json:"updated_at,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowKafkaRebalanceLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaRebalanceLogResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaRebalanceLogResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowKafkaLogTaskEntity struct {

	// **参数解释**： 日志记录ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 日志组ID。 **取值范围**： 不涉及。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释**： 日志流ID。 **取值范围**： 不涉及。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// **参数解释**： 仪表盘ID。 **取值范围**： 不涉及。
	DashboardId *string `json:"dashboard_id,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - OPEN：开启。 - CLOSE：关闭。 - CLOSING：关闭中。 - OPENING：开启中。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - REBALANCE：重平衡日志。 - topic_log：Topic日志。
	LogType *string `json:"log_type,omitempty"`

	// **参数解释**： 日志文件名。 **取值范围**： 不涉及。
	LogFileName *string `json:"log_file_name,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**： 更新时间。    **取值范围**： 不涉及。
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o ShowKafkaLogTaskEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaLogTaskEntity struct{}"
	}

	return strings.Join([]string{"ShowKafkaLogTaskEntity", string(data)}, " ")
}

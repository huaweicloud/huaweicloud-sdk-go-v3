package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisProgress 重分布进度信息
type RedisProgress struct {

	// **参数解释**： 已完成字节数。 **取值范围**： 不涉及。
	BytesDone *int64 `json:"bytes_done,omitempty"`

	// **参数解释**： 剩余字节数。 **取值范围**： 不涉及。
	ByteLeft *int64 `json:"byte_left,omitempty"`

	// **参数解释**： 完成表数量。 **取值范围**： 不涉及。
	TablesDone *int32 `json:"tables_done,omitempty"`

	// **参数解释**： 剩余表数量。 **取值范围**： 不涉及。
	TablesLeft *int32 `json:"tables_left,omitempty"`

	// **参数解释**： 表重分布进度。 **取值范围**： 1~100。
	TableProgress *int32 `json:"table_progress,omitempty"`

	// **参数解释**： 总进度。 **取值范围**： 1~100。
	TotalProgress *int32 `json:"total_progress,omitempty"`

	// **参数解释**： 重分布速度。 **取值范围**： 不涉及。
	RedisRate *string `json:"redis_rate,omitempty"`

	// **参数解释**： 预估时间。 **取值范围**： 不涉及。
	EstimatedTime *string `json:"estimated_time,omitempty"`

	// **参数解释**： 是否已完成。 **取值范围**： 不涉及。
	Completed *bool `json:"completed,omitempty"`

	// **参数解释**： 是否完成初始化。 **取值范围**： 不涉及。
	Initialed *bool `json:"initialed,omitempty"`

	// **参数解释**： 失败次数。 **取值范围**： 不涉及。
	FailCount *int32 `json:"fail_count,omitempty"`

	// **参数解释**： cm_ctl查询的重分布结果。 **取值范围**： 不涉及。
	Redistributing *bool `json:"redistributing,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 是否用户暂停。 **取值范围**： 不涉及。
	PauseByUser *bool `json:"pause_by_user,omitempty"`
}

func (o RedisProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisProgress struct{}"
	}

	return strings.Join([]string{"RedisProgress", string(data)}, " ")
}

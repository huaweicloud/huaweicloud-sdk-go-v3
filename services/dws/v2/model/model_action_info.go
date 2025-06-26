package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionInfo **参数解释**： 逻辑集群操作信息。 **取值范围**： 不涉及。
type ActionInfo struct {

	// **参数解释**： 操作名称。 **取值范围**： Create：创建逻辑集群 Expand：扩容逻辑集群 Restart：重启逻辑集群 Delete：删除逻辑集群 Shrink：缩容逻辑集群
	ActionName *string `json:"action_name,omitempty"`

	// **参数解释**： 操作进度，默认10。 **取值范围**： 0~100
	Progress *int32 `json:"progress,omitempty"`

	// **参数解释**： 操作是否完成。 **取值范围**： 不涉及。
	Completed *bool `json:"completed,omitempty"`

	// **参数解释**： 操作开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 操作结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 操作结果。默认为空字符串。 **取值范围**： success：成功。 failed：失败。
	Result *string `json:"result,omitempty"`

	// **参数解释**： 操作日志信息。 **取值范围**： 不涉及。
	Logs *string `json:"logs,omitempty"`
}

func (o ActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInfo struct{}"
	}

	return strings.Join([]string{"ActionInfo", string(data)}, " ")
}

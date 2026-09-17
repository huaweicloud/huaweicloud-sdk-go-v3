package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrecheckStatus **参数解释：** 升级前检查状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type PrecheckStatus struct {

	// **参数解释：** 状态 **约束限制：** 不涉及 **取值范围：** - Init：初始化 - Running：运行中 - Success：成功 - Failed：失败 - Error：错误  **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释：** 检查结果过期时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ExpireTimeStamp *string `json:"expireTimeStamp,omitempty"`

	// **参数解释：** 信息，一般是执行错误的日志信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Message *string `json:"message,omitempty"`

	ClusterCheckStatus *ClusterCheckStatus `json:"clusterCheckStatus,omitempty"`

	AddonCheckStatus *AddonCheckStatus `json:"addonCheckStatus,omitempty"`

	NodeCheckStatus *NodeCheckStatus `json:"nodeCheckStatus,omitempty"`
}

func (o PrecheckStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckStatus struct{}"
	}

	return strings.Join([]string{"PrecheckStatus", string(data)}, " ")
}

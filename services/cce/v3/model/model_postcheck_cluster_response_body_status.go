package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostcheckClusterResponseBodyStatus **参数解释：** 集群升级后确认的状态信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type PostcheckClusterResponseBodyStatus struct {

	// **参数解释：** 状态 **约束限制：** 不涉及 **取值范围：** - Success：成功 - Failed：失败 - Error：错误  **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`
}

func (o PostcheckClusterResponseBodyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostcheckClusterResponseBodyStatus struct{}"
	}

	return strings.Join([]string{"PostcheckClusterResponseBodyStatus", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigStatusConfigsStatus **参数解释**：各配置项的执行状态。 **约束限制**：仅在请求体/插件中存在对应配置项时返回。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigStatusConfigsStatus struct {
	ImageWarmUp *NodeconfigStatusConfigsStatusImageWarmUp `json:"image_warm_up,omitempty"`
}

func (o NodeconfigStatusConfigsStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigStatusConfigsStatus struct{}"
	}

	return strings.Join([]string{"NodeconfigStatusConfigsStatus", string(data)}, " ")
}

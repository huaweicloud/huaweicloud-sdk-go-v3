package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterUpgradeResponseAction struct {

	// **参数解释：** 当前集群版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释：** 目标集群版本，例如\"v1.23\" **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetVersion *string `json:"targetVersion,omitempty"`

	// **参数解释：** 目标集群的平台版本号，表示集群版本(version)下的内部版本，不支持用户指定。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetPlatformVersion *string `json:"targetPlatformVersion,omitempty"`

	Strategy *UpgradeStrategy `json:"strategy,omitempty"`

	// **参数解释：** 升级过程中指定的集群配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Config *interface{} `json:"config,omitempty"`
}

func (o ClusterUpgradeResponseAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterUpgradeResponseAction struct{}"
	}

	return strings.Join([]string{"ClusterUpgradeResponseAction", string(data)}, " ")
}

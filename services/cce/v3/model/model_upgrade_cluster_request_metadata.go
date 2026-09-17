package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeClusterRequestMetadata **参数解释：** 集群升级请求元数据。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeClusterRequestMetadata struct {

	// **参数解释：** API版本，固定值\"v3\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - v3  **默认取值：** v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释：** API类型，固定值\"UpgradeTask\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - UpgradeTask  **默认取值：** UpgradeTask
	Kind string `json:"kind"`
}

func (o UpgradeClusterRequestMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeClusterRequestMetadata struct{}"
	}

	return strings.Join([]string{"UpgradeClusterRequestMetadata", string(data)}, " ")
}

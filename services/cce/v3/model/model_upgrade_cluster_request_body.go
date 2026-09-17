package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeClusterRequestBody **参数解释：** 集群升级请求体。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeClusterRequestBody struct {
	Metadata *UpgradeClusterRequestMetadata `json:"metadata"`

	Spec *UpgradeSpec `json:"spec"`
}

func (o UpgradeClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeClusterRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeClusterRequestBody", string(data)}, " ")
}

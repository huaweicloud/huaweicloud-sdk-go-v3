package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterConfigurationsSpecPackages struct {

	// **参数解释：** 组件名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 组件配置项详情 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Configurations *[]ConfigurationItem `json:"configurations,omitempty"`
}

func (o ClusterConfigurationsSpecPackages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterConfigurationsSpecPackages struct{}"
	}

	return strings.Join([]string{"ClusterConfigurationsSpecPackages", string(data)}, " ")
}

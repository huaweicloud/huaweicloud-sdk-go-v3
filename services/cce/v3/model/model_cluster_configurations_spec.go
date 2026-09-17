package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterConfigurationsSpec **参数解释：** Configuration的规格信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type ClusterConfigurationsSpec struct {

	// **参数解释：** 组件配置项列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Packages []ClusterConfigurationsSpecPackages `json:"packages"`
}

func (o ClusterConfigurationsSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterConfigurationsSpec struct{}"
	}

	return strings.Join([]string{"ClusterConfigurationsSpec", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonCheckSpec 插件检查Spec-request结构体
type AddonCheckSpec struct {

	// **参数解释：** 集群ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterID string `json:"clusterID"`

	// **参数解释：** 插件检查信息列表，包含了需要检查的插件模板名称，插件实例ID，插件升级配置等。 **约束限制：** 不涉及
	AddonList []AddonInfo `json:"addonList"`
}

func (o AddonCheckSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonCheckSpec struct{}"
	}

	return strings.Join([]string{"AddonCheckSpec", string(data)}, " ")
}

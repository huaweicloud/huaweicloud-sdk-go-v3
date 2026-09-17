package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeAddonConfig **参数解释：** 升级时插件操作配置。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeAddonConfig struct {

	// **参数解释：** CCE插件名称 **约束限制：** 不涉及 **取值范围：** 集群中已安装的插件名称。[集群中已安装插件详情见[获取AddonInstance列表](https://support.huaweicloud.com/api-cce/cce_02_0326.html)](tag:hws) **默认取值：** 不涉及
	AddonTemplateName string `json:"addonTemplateName"`

	// **参数解释：** 升级插件的执行动作 **约束限制：** 不涉及 **取值范围：** - patch：表示升级插件版本  **默认取值：** 不涉及
	Operation string `json:"operation"`

	// **参数解释：** 目标插件版本号 **约束限制：** 目标插件版本必须与目标集群版本配套。[集群版本配套关系见[查询AddonTemplates列表](https://support.huaweicloud.com/api-cce/cce_02_0321.html)](tag:hws) **取值范围：** 不涉及 **默认取值：** 不涉及
	Version string `json:"version"`

	// **参数解释：** 插件参数列表，Key:Value格式。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Values *interface{} `json:"values,omitempty"`
}

func (o UpgradeAddonConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeAddonConfig struct{}"
	}

	return strings.Join([]string{"UpgradeAddonConfig", string(data)}, " ")
}

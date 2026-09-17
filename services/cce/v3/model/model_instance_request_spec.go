package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceRequestSpec **参数解释**： spec是集合类的元素类型，内容为插件实例安装/升级的具体请求信息。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type InstanceRequestSpec struct {

	// **参数解释**： 待安装、升级插件的版本号，例如1.0.0。 **约束限制**： - 安装：该参数非必传，如果不传，匹配集群支持的最新版本。 - 升级：该参数必传，需指定版本号。  **取值范围**： 不涉及 **默认取值**： 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： 集群ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClusterID string `json:"clusterID"`

	// **参数解释**： 插件模板安装参数（各插件不同），升级插件时需要填写全量安装参数，未填写参数将使用插件模板中的默认值，当前插件安装参数可通过查询插件实例接口获取。[安装参数请参考[插件实例字段说明](https://support.huaweicloud.com/api-cce/cce_02_0366.html)。](tag:hws)[安装参数请参考[插件实例字段说明](https://support.huaweicloud.com/intl/zh-cn/api-cce/cce_02_0366.html)。](tag:hws_hk) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Values map[string]interface{} `json:"values"`

	// **参数解释**： 待安装插件模板名称，如coredns。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AddonTemplateName string `json:"addonTemplateName"`
}

func (o InstanceRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRequestSpec struct{}"
	}

	return strings.Join([]string{"InstanceRequestSpec", string(data)}, " ")
}

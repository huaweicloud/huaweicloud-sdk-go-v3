package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationItem struct {

	// **参数解释：** 组件配置参数名称。  [当前集群支持的可配置组件及其参数详见[修改CCE集群配置](https://support.huaweicloud.com/usermanual-cce/cce_10_0213.html)，当前节点池支持的可配置组件及其参数详见[修改节点池配置](https://support.huaweicloud.com/usermanual-cce/cce_10_0652.html)。](tag:hws) [当前集群支持的可配置组件及其参数详见[修改CCE集群配置](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0213.html)，当前节点池支持的可配置组件及其参数详见[修改节点池配置](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0652.html)。](tag:hws_hk) **约束限制：** 若指定了不支持的组件或组件不支持的参数，该配置项将被忽略。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 组件配置参数值。  [当前集群支持的可配置组件及其参数详见[修改CCE集群配置](https://support.huaweicloud.com/usermanual-cce/cce_10_0213.html)，当前节点池支持的可配置组件及其参数详见[修改节点池配置](https://support.huaweicloud.com/usermanual-cce/cce_10_0652.html)。](tag:hws) [当前集群支持的可配置组件及其参数详见[修改CCE集群配置](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0213.html)，当前节点池支持的可配置组件及其参数详见[修改节点池配置](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0652.html)。](tag:hws_hk) **约束限制：** 若指定了不支持的组件或组件不支持的参数，该配置项将被忽略。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Value *interface{} `json:"value,omitempty"`
}

func (o ConfigurationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationItem struct{}"
	}

	return strings.Join([]string{"ConfigurationItem", string(data)}, " ")
}

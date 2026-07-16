package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolSpecUpdate
type NodePoolSpecUpdate struct {
	NodeTemplate *NodeSpecUpdate `json:"nodeTemplate,omitempty"`

	// **参数解释：** 节点池期望节点个数。 **约束限制：** 更新节点池时，此字段为必填字段。 > 注意：如果更新节点池时不填此字段，节点池期望节点个数将取默认值0，如果此时节点池节点个数大于0将导致节点池缩容。  **取值范围：** 大于0，小于集群节点规模。 **默认取值：** 0
	InitialNodeCount int32 `json:"initialNodeCount"`

	// **参数解释：** 该参数用于控制更新节点池时 **节点池期望节点个数(spec.initialNodeCount)** 的默认行为。当该参数未设置或者为false时，如果用户请求Body体中未设置spec.initialNodeCount，更新时将自动初始化spec.initialNodeCount为0。当该参数为true时，将忽略spec.initialNodeCount参数。 > 当用户不需要更新节点池spec.initialNodeCount时，必须显示的设置该参数为true，同时在更新节点池Body体中不设置spec.initialNodeCount。  **约束限制：** 不涉及 **取值范围：** - false：更新节点池时，如果spec.initialNodeCount参数未设置，将初始化spec.initialNodeCount为0。 > 如果节点池当前spec.initialNodeCount 不等于0将导致节点池缩容。  - true：更新节点池时，忽略spec.initialNodeCount参数，节点池spec.initialNodeCount参数将保持原样。  **默认取值：** false
	IgnoreInitialNodeCount *bool `json:"ignoreInitialNodeCount,omitempty"`

	Autoscaling *NodePoolNodeAutoscaling `json:"autoscaling,omitempty"`

	NodeManagementUpdate *NodeManagement `json:"nodeManagementUpdate,omitempty"`

	// 节点池自定义安全组相关配置。支持节点池新扩容节点绑定指定的安全组。  - 未指定安全组ID，新建节点将添加Node节点默认安全组。  - 指定有效安全组ID，新建节点将使用指定安全组。  - 指定安全组，应避免对CCE运行依赖的端口规则进行修改。[详细设置请参考[集群安全组规则配置](https://support.huaweicloud.com/cce_faq/cce_faq_00265.html)。](tag:hws)[详细设置请参考[集群安全组规则配置](https://support.huaweicloud.com/intl/zh-cn/cce_faq/cce_faq_00265.html)。](tag:hws_hk)
	CustomSecurityGroups *[]string `json:"customSecurityGroups,omitempty"`

	// **参数解释：** 是否同步K8S污点。 **约束限制**： 不涉及 **取值范围：** - 填写为refresh，K8S污点的改动将会被同步更新到存量节点上。 - 填写为ignore，节点池K8S污点将不会同步更新到存量节点上。  **默认取值：** 无
	TaintPolicyOnExistingNodes *string `json:"taintPolicyOnExistingNodes,omitempty"`

	// **参数解释：** 是否同步K8S标签。 **约束限制**： 不涉及 **取值范围：** - 填写为refresh，K8S标签的改动将会被同步更新到存量节点上。 - 填写为ignore，K8S标签将不会同步更新到存量节点上。  **默认取值：** 无
	LabelPolicyOnExistingNodes *string `json:"labelPolicyOnExistingNodes,omitempty"`

	// **参数解释：** 是否同步节点池资源标签。 **约束限制**： 不涉及 **取值范围：** - 填写为refresh，节点池资源标签标签的改动将会被同步更新到存量节点上。 - 填写为ignore，节点池资源标签标签将不会同步更新到存量节点上。  **默认取值：** 无
	UserTagsPolicyOnExistingNodes *string `json:"userTagsPolicyOnExistingNodes,omitempty"`

	// 节点池扩展伸缩组配置列表，详情参见ExtensionScaleGroup类型定义
	ExtensionScaleGroups *[]ExtensionScaleGroup `json:"extensionScaleGroups,omitempty"`
}

func (o NodePoolSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolSpecUpdate struct{}"
	}

	return strings.Join([]string{"NodePoolSpecUpdate", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterUpgradeAction **参数解释：** 集群升级动作定义，包含目标版本、升级策略、插件配置等。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type ClusterUpgradeAction struct {

	// **参数解释：** 插件配置列表，CCE会在集群升级过程中按照配置对插件进行升级 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Addons *[]UpgradeAddonConfig `json:"addons,omitempty"`

	// **参数解释：** 节点池内节点升级顺序配置。key表示节点池ID，默认节点池取值为\"DefaultPool\" **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NodeOrder map[string][]NodePriority `json:"nodeOrder,omitempty"`

	// **参数解释：** 节点池升级顺序配置，key/value对格式。key表示节点池ID，默认节点池取值为\"DefaultPool\"，value表示对应节点池的优先级，默认值为0，优先级最低，数值越大优先级越高 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NodePoolOrder map[string]int32 `json:"nodePoolOrder,omitempty"`

	Strategy *UpgradeStrategy `json:"strategy"`

	// **参数解释：** 升级的目标集群版本，例如\"v1.23\" **约束限制：** 只能升级到高版本，不允许填写等于或低于当前集群版本的值 **取值范围：** CCE支持的集群版本 **默认取值：** 不涉及
	TargetVersion string `json:"targetVersion"`

	// **参数解释：** 是否在集群升级流程中跳过升级前检查。 **约束限制：** 不涉及 **取值范围：** - false：表示在集群升级流程中会执行升级前检查。 - true：表示在集群升级流程中跳过升级前检查。  **默认取值：** false
	IsOnlyUpgrade *bool `json:"isOnlyUpgrade,omitempty"`

	// **参数解释：** 指定集群使用的委托。该委托用于生成集群中组件使用的临时访问凭证，在集群中自动创建其他相关云服务的资源时会使用该委托权限。 当不传时，集群将优先继承原有配置，若原先未配置，则自动选择使用CCE的默认委托CCEAutoClusterAgency；当传空时，自动选择使用CCE的默认委托CCEAutoClusterAgency。  [ > 关于CCE系统委托的说明详情参见[系统委托说明](https://support.huaweicloud.com/usermanual-cce/cce_10_0556.html)](tag:hws) [ > 关于CCE系统委托的说明详情参见[系统委托说明](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0556.html)](tag:hws_hk)  **约束限制：** 仅v1.28.15-r90、v1.29.15-r50、v1.30.14-r50、v1.31.14-r10、v1.32.9-r10、v1.33.7-r10、v1.34.3-r0及以上版本集群支持该参数 **取值范围：** 不涉及 **默认取值：** 空
	AgencyName *string `json:"agencyName,omitempty"`
}

func (o ClusterUpgradeAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterUpgradeAction struct{}"
	}

	return strings.Join([]string{"ClusterUpgradeAction", string(data)}, " ")
}

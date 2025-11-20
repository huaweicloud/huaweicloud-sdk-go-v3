package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInformationSpec
type ClusterInformationSpec struct {

	// **参数解释：** 指定集群使用的委托。该委托用于生成集群中组件使用的临时访问凭证，在集群中自动创建其他相关云服务的资源时会使用该委托权限。当不传或为空时，集群将自动选择使用CCE的系统委托cce_admin_trust或cce_cluster_agency。  [ > 关于CCE系统委托的说明详情参见[系统委托说明](https://support.huaweicloud.com/usermanual-cce/cce_10_0556.html)](tag:hws) [ > 关于CCE系统委托的说明详情参见[系统委托说明](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0556.html)](tag:hws_hk)  **约束限制：** 仅1.27及以上版本集群支持该参数  **取值范围：** 不涉及 **默认取值：** 空
	AgencyName *string `json:"agencyName,omitempty"`

	// **参数解释：** 集群的描述信息。 **约束限制：** 仅运行和扩容状态（Available、ScalingUp、ScalingDown）的集群允许修改。 **取值范围：** 字符取值范围[0,200]。不包含~$%^&*<>[]{}()'\"#\\等特殊字符。 **默认取值：** 无
	Description *string `json:"description,omitempty"`

	// 集群的API Server服务端证书中的自定义SAN（Subject Alternative Name）字段，遵从SSL标准X509定义的格式规范。  1. 不允许出现同名重复。 2. 格式符合IP和域名格式。  示例: ``` SAN 1: DNS Name=example.com SAN 2: DNS Name=www.example.com SAN 3: DNS Name=example.net SAN 4: IP Address=93.184.216.34 ```
	CustomSan *[]string `json:"customSan,omitempty"`

	ContainerNetwork *ContainerNetworkUpdate `json:"containerNetwork,omitempty"`

	EniNetwork *EniNetworkUpdate `json:"eniNetwork,omitempty"`

	HostNetwork *ClusterInformationSpecHostNetwork `json:"hostNetwork,omitempty"`

	// **参数解释：** 集群删除保护，如果开启后用户将无法删除该集群。 **约束限制：** 不涉及 **取值范围：** - true: 开启集群删除保护 - false: 关闭集群删除保护  **默认取值：** 默认false
	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	// **参数解释：** 配置参数后，配置CCE Standard/Turbo集群是否启用自动升配功能。 **约束限制：** 当前集群自动升配功能受限开放。 集群支持范围： - 版本范围：v1.27及以上 - 规格范围：cce.s2.*规格的集群支持启用自动升配  **取值范围：** - true: 启用自动升配能力 - false: 禁用自动升配能力  **默认取值：** 默认false，未指定则不更新此参数
	EnableAutoResizing *bool `json:"enableAutoResizing,omitempty"`
}

func (o ClusterInformationSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInformationSpec struct{}"
	}

	return strings.Join([]string{"ClusterInformationSpec", string(data)}, " ")
}

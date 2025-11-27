package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperProtectionConfigsRequest Request Object
type ListWebTamperProtectionConfigsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 网页防篡改类型 **约束限制**: 不涉及。 **取值范围**: - container_wtp：容器网页防篡改  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 网页防篡改配置ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 网站应用的名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	WebAppName *string `json:"web_app_name,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-512位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 防护类型 **约束限制**: 不涉及。 **取值范围**: - cluster：集群范围内防护 - host：主机范围内防护  **默认取值**: 不涉及
	ProtectType *string `json:"protect_type,omitempty"`

	// **参数解释**: 防护状态 **约束限制**: 不涉及。 **取值范围**: - not_protect：未防护 - protected：防护中 - partial_protected：部分防护 - protect_failed：防护失败 - redundant：防护冗余  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 集群名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群id **约束限制**： 不涉及 **取值范围**： 字符长度0-128位 **默认取值**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机私有IP **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostPrivateIp *string `json:"host_private_ip,omitempty"`
}

func (o ListWebTamperProtectionConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperProtectionConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListWebTamperProtectionConfigsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonHostRequest Request Object
type ListCommonHostRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 必填 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **约束限制**: 不涉及 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位） **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 策略名称 **约束限制**: 不涉及 **取值范围**: **取值范围**: - av_detect_feature：AV策略  **默认取值**: 无
	FeatureName *string `json:"feature_name,omitempty"`

	// **参数解释**: 服务器组的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： Agent状态 **约束限制**: 不涉及 **取值范围**： - installed：已安装 - not_installed：未安装 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中  **默认取值**: 不涉及
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 集群ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 集群名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 主机开通的版本,高于该版本 **约束限制**: 不涉及 **取值范围**： - hss.version.basic：基础版。 - hss.version.advanced：专业版。 - hss.version.enterprise：企业版。 - hss.version.premium：旗舰版。 - hss.version.wtp：网页防篡改版。 - hss.version.container.enterprise：容器版。  **默认取值**: 不涉及
	VersionNameUpper *string `json:"version_name_upper,omitempty"`

	// **参数解释**： 主机开通的版本,低于该版本 **约束限制**: 不涉及 **取值范围**： - hss.version.basic：基础版。 - hss.version.advanced：专业版。 - hss.version.enterprise：企业版。 - hss.version.premium：旗舰版。 - hss.version.wtp：网页防篡改版。 - hss.version.container.enterprise：容器版。  **默认取值**: 不涉及
	VersionNameLower *string `json:"version_name_lower,omitempty"`
}

func (o ListCommonHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonHostRequest struct{}"
	}

	return strings.Join([]string{"ListCommonHostRequest", string(data)}, " ")
}

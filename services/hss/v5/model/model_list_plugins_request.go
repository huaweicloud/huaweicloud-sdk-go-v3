package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginsRequest Request Object
type ListPluginsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 客户端状态 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**： 检测结果 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释**： 主机状态 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**： 操作系统类型 **约束限制**： 不涉及 **取值范围**： - Linux：Linux操作系统。 - Windows：Windows操作系统。  **默认取值**： 不涉及
	OsType *string `json:"os_type,omitempty"`

	// 公网或私网IP
	IpAddr *string `json:"ip_addr,omitempty"`

	// 防护状态
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`

	// 策略组ID
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// 策略组名称
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// 资产标签
	Label *string `json:"label,omitempty"`

	// 收费模式
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**： 是否强制从ECS同步主机 **约束限制**： 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**： false
	Refresh *bool `json:"refresh,omitempty"`

	// 是否返回比当前版本高的所有版本
	AboveVersion *bool `json:"above_version,omitempty"`

	// **参数解释**： 插件名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位  **默认取值**： opa-docker-authz
	Name string `json:"name"`

	// **参数解释**： 主机开通的版本 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件类型 **约束限制**： 不涉及 **取值范围**： - opa-docker-authz：docker插件。  **默认取值**： opa-docker-authz
	Plugin *string `json:"plugin,omitempty"`

	// **参数解释**： 是否非华为云 **约束限制**： 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**： false
	OutsideHost *bool `json:"outside_host,omitempty"`
}

func (o ListPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListPluginsRequest", string(data)}, " ")
}

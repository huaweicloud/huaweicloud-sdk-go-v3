package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventTypeRequest Request Object
type ListEventTypeRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 开始时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**: 结束时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// 查询时间范围天数，与自定义查询时间begin_time，end_time互斥
	LastDays *int32 `json:"last_days,omitempty"`

	// **参数解释**: 事件类别 **约束限制**: 不涉及 **取值范围**: - host : 主机安全事件 - container : 容器安全事件 - serverless : serverless场景安全事件 **默认取值**: 不涉及
	Category string `json:"category"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 容器实例名称
	ContainerName *string `json:"container_name,omitempty"`

	// 处置状态，包含如下:   - unhandled ：未处理   - handled : 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 威胁等级，包含如下:   - Security ：安全   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 危急
	Severity *string `json:"severity,omitempty"`

	// 威胁等级，包含如下:   - Security ：安全   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 危急
	SeverityList *[]string `json:"severity_list,omitempty"`

	// 攻击标识，包含如下：   - attack_success : 攻击成功   - attack_attempt : 攻击尝试   - attack_blocked : 攻击被阻断   - abnormal_behavior : 异常行为   - collapsible_host : 主机失陷   - system_vulnerability : 系统脆弱性
	AttackTag *string `json:"attack_tag,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 事件标签列表，例如:[\"热点事件\"]
	TagList *[]string `json:"tag_list,omitempty"`

	// ATT&CK攻击阶，包含如下：   - Reconnaissance : 侦察   - Initial Access : 初始访问   - Execution : 执行   - Persistence : 持久化   - Privilege Escalation : 权限提升   - Defense Evasion : 防御绕过   - Credential Access : 凭据访问   - Command and Control : 命令与控制   - Impact : 影响破坏
	AttCk *string `json:"att_ck,omitempty"`

	// 告警名称
	EventName *string `json:"event_name,omitempty"`
}

func (o ListEventTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventTypeRequest struct{}"
	}

	return strings.Join([]string{"ListEventTypeRequest", string(data)}, " ")
}

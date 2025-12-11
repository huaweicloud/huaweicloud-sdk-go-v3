package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventAttCkRequest Request Object
type ListEventAttCkRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Region string `json:"region"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 查询时间范围天数，与自定义查询时间begin_time，end_time互斥。 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值30 **默认取值**: 不涉及
	LastDays *int32 `json:"last_days,omitempty"`

	// **参数解释**: 事件类别 **约束限制**: 不涉及 **取值范围**: - host : 主机安全事件 - container : 容器安全事件 - serverless : serverless场景安全事件  **默认取值**: 不涉及
	Category string `json:"category"`

	// **参数解释**： 服务器名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器ID **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 服务器私有IP **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 服务器公网IP **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 容器实例名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-512位 **默认取值**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 事件类型 **约束限制**: 不涉及 **取值范围**: - 1001：通用恶意软件 - 1002：病毒 - 1003：蠕虫 - 1004：木马 - 1005：僵尸网络 - 1006：后门 - 1010：Rootkit - 1011：勒索软件 - 1012：黑客工具 - 1015：Webshell - 1016：挖矿 - 1017：反弹Shell - 2001：一般漏洞利用 - 2012：远程代码执行 - 2047：Redis漏洞利用 - 2048：Hadoop漏洞利用 - 2049：MySQL漏洞利用 - 3002：文件提权 - 3003：进程提权 - 3004：关键文件变更 - 3005：文件/目录变更 - 3007：进程异常行为 - 3015：高危命令执行 - 3018：异常Shell - 3026：crontab提权 - 3027：Crontab可疑任务 - 3029：系统安全防护被禁用 - 3030：备份删除 - 3031：异常注册表操作 - 3036：容器镜像阻断 - 4002：暴力破解 - 4004：异常登录 - 4006：非法系统账号 - 4014：用户账号添加 - 4020：用户密码窃取 - 6002：端口扫描 - 6003：主机扫描 - 13001：Kubernetes事件删除 - 13002：Pod异常行为 - 13003：枚举用户信息 - 13004：绑定集群用户角色  **默认取值**: 不涉及
	EventType *int32 `json:"event_type,omitempty"`

	// **参数解释**: 处置状态 **约束限制**: 不涉及 **取值范围**: - unhandled：未处理 - handled：已处理  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 威胁等级 **约束限制**: 不涉及 **取值范围**: - Security：安全 - Low：低危 - Medium：中危 - High：高危 - Critical：危急  **默认取值**: 不涉及
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 威胁等级 **约束限制**: 不涉及 **取值范围**: - Security：安全 - Low：低危 - Medium：中危 - High：高危 - Critical：危急  **默认取值**: 不涉及
	SeverityList *[]string `json:"severity_list,omitempty"`

	// **参数解释**: 攻击标识 **约束限制**: 不涉及 **取值范围**: - attack_success：攻击成功 - attack_attempt：攻击尝试 - attack_blocked：攻击被阻断 - abnormal_behavior：异常行为 - collapsible_host：主机失陷 - system_vulnerability：系统脆弱性  **默认取值**: 不涉及
	AttackTag *string `json:"attack_tag,omitempty"`

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// 事件标签列表，例如:[\"热点事件\"]
	TagList *[]string `json:"tag_list,omitempty"`

	// **参数解释**： 告警名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	EventName *string `json:"event_name,omitempty"`
}

func (o ListEventAttCkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventAttCkRequest struct{}"
	}

	return strings.Join([]string{"ListEventAttCkRequest", string(data)}, " ")
}

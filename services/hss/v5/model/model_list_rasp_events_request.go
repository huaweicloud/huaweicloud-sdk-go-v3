package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRaspEventsRequest Request Object
type ListRaspEventsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// Host Id
	HostId string `json:"host_id"`

	// **参数解释**: 应用防护事件的查询起始时间（Unix时间戳），与end_time配合筛选指定时间段内的事件 **时间格式** Unix时间戳（精确到毫秒，如1736414463000表示2024-12-10 10:41:03） **约束限制**: 需小于end_time，否则返回空结果；时间戳需为有效时间（1970-01-01 00:00:00至今） **取值范围**: 取值0-9223372036854775807 **默认取值**: 无
	StartTime int64 `json:"start_time"`

	// **参数解释**: 查询时间段的终止时间，毫秒级时间戳(ms) **约束限制**: 不涉及 **取值范围**: 取值0-9223372036854775807 **默认取值**: 无
	EndTime int64 `json:"end_time"`

	// **参数解释** 应用防护的应用类型，用于筛选指定类型应用的防护事件 **约束限制** 当前仅支持java类型，传其他值返回空结果，区分大小写 **取值范围** - java：Java语言开发的应用防护事件 **默认取值** 无（查询所有支持的应用类型事件）
	AppType *string `json:"app_type,omitempty"`

	// **参数解释** 应用防护事件的告警级别，用于筛选指定严重程度的事件 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - Security：信息级 - Low：低危 - Medium：中危 - High：高危 - Critical：紧急 **默认取值** 无
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 应用防护事件的攻击类型标识，用于筛选指定攻击类型的事件 **约束限制** 取值区分大小写，必须与指定格式一致，否则返回空结果 **取值范围** - Attack Success：攻击成功 - Attack Attempt：攻击尝试 - Attack Blocked：攻击被阻断 - Abnormal Behavior：异常行为 - Collapsible Host：主机失陷 - System Vulnerability：系统脆弱性 **默认取值** 无
	AttackTag *string `json:"attack_tag,omitempty"`

	// **参数解释** 应用防护的启用状态，用于筛选指定防护状态下的事件 **约束限制** 取值区分大小写，必须在指定范围内，否则返回空结果 **取值范围** - closed：未开启防护 - opened：已开启防护 **默认取值** 无（查询所有防护状态的事件）
	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ListRaspEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRaspEventsRequest struct{}"
	}

	return strings.Join([]string{"ListRaspEventsRequest", string(data)}, " ")
}

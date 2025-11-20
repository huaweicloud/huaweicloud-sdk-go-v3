package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAlertNoticeConfigBody struct {

	// **参数解释：** 功能启用状态：true表示启用当前配置，false表示禁用。 **约束限制：** 不涉及 **取值范围：** - true - false  **默认取值：** true
	Enabled bool `json:"enabled"`

	// **参数解释：** 预留参数，默认为false **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** false
	PreferHtml bool `json:"prefer_html"`

	// **参数解释：** 通知模板，关联用于发送告警通知的SMN主题URN,通过“消息通知服务”获取 查询可使用的主题，通过 云日志服务的“查询SMN主题”接口，返回体中的\"topic_urn\"字段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopicUrn string `json:"topic_urn"`

	// **参数解释：** 通间间隔,单位为分钟。当通知类型为防护事件时,该参数表示在该时间间隔内,攻击次数等于或者大于设定阈值时,将发送告警通知,支持的值:5、15、30、60、120、360、720、1440;当通知类型为证书到期时,该参数表示每隔多长时间发送一次告警通知,支持的值为1440、10080(单位为分钟)。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sendfreq *int32 `json:"sendfreq,omitempty"`

	// **参数解释：** 地区，指定告警通知的语言或地域相关配置 **约束限制：** 不涉及 **取值范围：** - zh-cn - en-us  **默认取值：** 不涉及
	Local string `json:"local"`

	// **参数解释：** 当通知类型为防护事件时,需要填写该参数。在该时间间隔内,当攻击次数大于或等于您设置的阈值时才会发送告警通知 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Times *int32 `json:"times,omitempty"`

	// **参数解释：** 告警名称，用于标识告警通知配置的名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 告警类型，区分不同场景的告警（如防护规则触发、资源异常等） **约束限制：** 不涉及 **取值范围：** - threat_alert_notice：威胁告警通知  - cert_alert_notice：证书告警通知, - rule_alert_notice：规则告警通知, - cname_ip_alert_notice：域名 / IP 告警通知  **默认取值：** 不涉及
	NoticeClass string `json:"notice_class"`

	// **参数解释：** 提前通知天数，针对过期类告警提前发送通知的天数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NearlyExpiredTime *int64 `json:"nearly_expired_time,omitempty"`

	// **参数解释：** 是否所有企业项目，标识该告警配置是否适用于所有企业项目 **约束限制：** 不涉及 **取值范围：** - true - false  **默认取值：** true
	IsAllEnterpriseProject bool `json:"is_all_enterprise_project"`

	// **参数解释：** 描述，对告警通知配置的补充说明 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 威胁类型范围：指定需要告警的威胁类型，如[\"all\"]表示所有威胁 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** all
	Threat *[]string `json:"threat,omitempty"`

	// **参数解释：** 指定需要告警的规则类型，[\"all\"]表示所有规则类型。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** all
	RuleType *[]string `json:"rule_type,omitempty"`
}

func (o CreateAlertNoticeConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertNoticeConfigBody struct{}"
	}

	return strings.Join([]string{"CreateAlertNoticeConfigBody", string(data)}, " ")
}

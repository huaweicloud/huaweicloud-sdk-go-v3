package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertNoticeConfig struct {

	// **参数解释：** 告警id，用于唯一标识一条告警通知配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 是否开启告警，控制该告警通知配置的启用/禁用状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释：** 通知模板，关联用于发送告警通知的SMN主题URN 查询可使用的主题，通过 云日志服务的“查询SMN主题”接口，返回体中的\"topic_urn\"字段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释：** 通知频率，控制告警通知的发送间隔 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sendfreq *int32 `json:"sendfreq,omitempty"`

	// **参数解释：** 地区，指定告警通知的语言或地域相关配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Local *string `json:"local,omitempty"`

	// **参数解释：** 通知频率（补充说明，与sendfreq协同控制告警发送频次） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Times *int32 `json:"times,omitempty"`

	// **参数解释：** 告警名称，用于标识告警通知配置的名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 告警类型，区分不同场景的告警（如防护规则触发、资源异常等） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NoticeClass *string `json:"notice_class,omitempty"`

	// **参数解释：** 提前通知天数，针对过期类告警提前发送通知的天数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NearlyExpiredTime *int64 `json:"nearly_expired_time,omitempty"`

	// **参数解释：** 是否所有企业项目，标识该告警配置是否适用于所有企业项目 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IsAllEnterpriseProject *bool `json:"is_all_enterprise_project,omitempty"`

	// **参数解释：** 描述，对告警通知配置的补充说明 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// 预留参数，可忽略
	PreferHtml *bool `json:"prefer_html,omitempty"`

	// 告警的事件类型
	Threat *[]string `json:"threat,omitempty"`
}

func (o AlertNoticeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertNoticeConfig struct{}"
	}

	return strings.Join([]string{"AlertNoticeConfig", string(data)}, " ")
}

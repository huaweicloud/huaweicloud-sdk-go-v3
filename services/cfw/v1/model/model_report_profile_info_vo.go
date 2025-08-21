package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportProfileInfoVo struct {

	// **参数解释**： 报告类型 **取值范围**： daily 日报 weekly 周报 custom 自定义报告
	Category *string `json:"category,omitempty"`

	// **参数解释**： 描述信息 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 模板名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 发送时间 **取值范围**： 不涉及
	SendPeriod *int32 `json:"send_period,omitempty"`

	// **参数解释**： 发送星期 **取值范围**： 不涉及
	SendWeekDay *int32 `json:"send_week_day,omitempty"`

	StatisticPeriod *StatisticPeriod `json:"statistic_period,omitempty"`

	// **参数解释**： 启用状态 **取值范围**：  0 关闭 1 启用
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 通知主题名称 **取值范围**：  不涉及
	TopicName *string `json:"topic_name,omitempty"`

	// **参数解释**： 通知主题urn **取值范围**：  不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**： 通知方式 **取值范围**： 0 SMN通知方式 1 不需要通知
	SubscriptionType *int32 `json:"subscription_type,omitempty"`
}

func (o ReportProfileInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportProfileInfoVo struct{}"
	}

	return strings.Join([]string{"ReportProfileInfoVo", string(data)}, " ")
}

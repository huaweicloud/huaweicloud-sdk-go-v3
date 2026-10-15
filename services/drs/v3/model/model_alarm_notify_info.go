package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmNotifyInfo 异常通知信息。
type AlarmNotifyInfo struct {

	// **参数解释**：  订阅延迟时间，即异常告警通知的延迟时间，单位为秒。  **约束限制**：  不涉及。  **取值范围**：  1-3600。  **默认取值**：  0
	DelayTime *int64 `json:"delay_time,omitempty"`

	// **参数解释**：  RTO（Recovery Time Objective，恢复时间目标）延迟时间，即容灾恢复时间目标的延迟阈值，单位为秒。  **约束限制**：  不涉及。  **取值范围**：  1-3600。  **默认取值**：  0
	RtoDelay *int64 `json:"rto_delay,omitempty"`

	// **参数解释**：  RPO（Recovery Point Objective，恢复点目标）延迟时间，即容灾恢复点目标的延迟阈值，单位为秒。  **约束限制**：  不涉及。  **取值范围**：  1-3600。  **默认取值**：  0
	RpoDelay *int64 `json:"rpo_delay,omitempty"`

	// **参数解释**：  异常告警是否通知用户。  **约束限制**：  不涉及。  **取值范围**：  - true：通知用户 - false：不通知用户  **默认取值**：  false
	AlarmToUser *bool `json:"alarm_to_user,omitempty"`

	// **参数解释**：  收件方式与信息体列表，用于配置异常告警的通知渠道，最多填写两个，详情请参见SubscriptionInfo对象定义。  **约束限制**：  - 最多填写两个。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty"`
}

func (o AlarmNotifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmNotifyInfo struct{}"
	}

	return strings.Join([]string{"AlarmNotifyInfo", string(data)}, " ")
}

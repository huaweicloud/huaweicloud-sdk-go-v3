package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrackerConfigBody TrackerConfig对象
type TrackerConfigBody struct {
	Channel *ChannelConfigBody `json:"channel"`

	Selector *SelectorConfigBody `json:"selector"`

	// 存储历史信息的天数
	RetentionPeriodInDays *int32 `json:"retention_period_in_days,omitempty"`

	// IAM委托名称
	AgencyName string `json:"agency_name"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	FrozenStatus *FrozenStatus `json:"frozen_status,omitempty"`
}

func (o TrackerConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerConfigBody struct{}"
	}

	return strings.Join([]string{"TrackerConfigBody", string(data)}, " ")
}

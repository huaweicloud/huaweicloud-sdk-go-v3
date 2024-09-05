package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrackerConfigResponse Response Object
type ShowTrackerConfigResponse struct {
	Channel *ChannelConfigBody `json:"channel,omitempty"`

	Selector *SelectorConfigBody `json:"selector,omitempty"`

	// 存储历史信息的天数
	RetentionPeriodInDays *int32 `json:"retention_period_in_days,omitempty"`

	// IAM委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	FrozenStatus   *FrozenStatus `json:"frozen_status,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowTrackerConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrackerConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowTrackerConfigResponse", string(data)}, " ")
}

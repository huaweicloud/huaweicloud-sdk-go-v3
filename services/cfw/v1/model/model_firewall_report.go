package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirewallReport struct {
	AttackInfo *AttackReport `json:"attack_info,omitempty"`

	// **参数解释**： 报告类型 **取值范围**： 不涉及
	Category *string `json:"category,omitempty"`

	InternetFirewall *InternetReport `json:"internet_firewall,omitempty"`

	// **参数解释**： 发送时间 **取值范围**： 不涉及
	SendTime *int64 `json:"send_time,omitempty"`

	StatisticPeriod *StatisticPeriod `json:"statistic_period,omitempty"`

	VpcFirewall *VpcReport `json:"vpc_firewall,omitempty"`
}

func (o FirewallReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallReport struct{}"
	}

	return strings.Join([]string{"FirewallReport", string(data)}, " ")
}

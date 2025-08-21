package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Overview struct {
	AccessPolicies *AccessPolicy `json:"access_policies,omitempty"`

	Assets *ChangedVo `json:"assets,omitempty"`

	AttackEvent *AttackEvent `json:"attack_event,omitempty"`

	TrafficPeak *TrendVo `json:"traffic_peak,omitempty"`
}

func (o Overview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Overview struct{}"
	}

	return strings.Join([]string{"Overview", string(data)}, " ")
}

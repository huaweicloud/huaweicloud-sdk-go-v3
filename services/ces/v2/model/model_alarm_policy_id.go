package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmPolicyId 告警策略ID。
type AlarmPolicyId struct {
}

func (o AlarmPolicyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmPolicyId struct{}"
	}

	return strings.Join([]string{"AlarmPolicyId", string(data)}, " ")
}

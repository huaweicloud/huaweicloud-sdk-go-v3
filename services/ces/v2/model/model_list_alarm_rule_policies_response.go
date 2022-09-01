package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAlarmRulePoliciesResponse struct {

	// 策略信息
	Policies *[]Policy `json:"policies,omitempty" xml:"policies"`

	// 指定告警规则对应的策略总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulePoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulePoliciesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicySwitchStatusResponse Response Object
type ShowPolicySwitchStatusResponse struct {

	// **参数解释**： 策略名称 **取值范围**： 字符长度1-128位
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 策略总开关是否开启 **取值范围**： - true：是。 - false：否。
	Enable         *bool `json:"enable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPolicySwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicySwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicySwitchStatusResponse", string(data)}, " ")
}

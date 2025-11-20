package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageBaselineStatisticResponse Response Object
type ShowImageBaselineStatisticResponse struct {

	// **参数解释**: 弱口令检测 **取值范围**: 最小值0，最大值2147483647
	ImageWeakPwd *int32 `json:"image_weak_pwd,omitempty"`

	// **参数解释**: 口令复杂度策略检测 **取值范围**: 最小值0，最大值2147483647
	PwdPolicy *int32 `json:"pwd_policy,omitempty"`

	// **参数解释**: 服务器配置检测 **取值范围**: 最小值0，最大值2147483647
	SecurityCheck  *int32 `json:"security_check,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageBaselineStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageBaselineStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowImageBaselineStatisticResponse", string(data)}, " ")
}

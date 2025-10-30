package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineStatisticResponse Response Object
type ShowBaselineStatisticResponse struct {

	// **参数解释**: 弱口令检测数量 **取值范围**: 取值0-9223372036854775807
	HostWeakPwd *int64 `json:"host_weak_pwd,omitempty"`

	// **参数解释**: 口令复杂度策略检测数量 **取值范围**: 取值0-9223372036854775807
	PwdPolicy *int64 `json:"pwd_policy,omitempty"`

	// **参数解释**: 服务器配置检测数量 **取值范围**: 取值0-9223372036854775807
	SecurityCheck  *int64 `json:"security_check,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowBaselineStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowBaselineStatisticResponse", string(data)}, " ")
}

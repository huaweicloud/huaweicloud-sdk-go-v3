package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTimerRuleReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“TimerRule”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	Spec *UpdateTimerRuleDetails `json:"spec,omitempty"`
}

func (o CreateTimerRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTimerRuleReq struct{}"
	}

	return strings.Join([]string{"CreateTimerRuleReq", string(data)}, " ")
}

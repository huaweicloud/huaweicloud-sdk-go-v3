package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTimerRuleReq struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *TimeRuleKindObj `json:"kind,omitempty"`

	Spec *UpdateTimerRuleDetails `json:"spec,omitempty"`
}

func (o CreateTimerRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTimerRuleReq struct{}"
	}

	return strings.Join([]string{"CreateTimerRuleReq", string(data)}, " ")
}

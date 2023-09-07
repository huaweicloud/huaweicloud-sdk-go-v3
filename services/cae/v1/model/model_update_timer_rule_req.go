package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTimerRuleReq struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *TimeRuleKindObj `json:"kind"`

	Spec *UpdateTimerRuleDetails `json:"spec"`
}

func (o UpdateTimerRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimerRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateTimerRuleReq", string(data)}, " ")
}

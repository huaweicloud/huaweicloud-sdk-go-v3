package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTimerRuleReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// 资源种类。
	Kind string `json:"kind"`

	Spec *UpdateTimerRuleDetails `json:"spec"`
}

func (o UpdateTimerRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimerRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateTimerRuleReq", string(data)}, " ")
}

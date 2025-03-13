package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradePlan struct {

	// API类型，固定值“UpgradePlan”，该值不可修改。
	Kind string `json:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion string `json:"apiVersion"`

	Metadata *Metadata `json:"metadata"`

	Spec *UpgradePlanSpec `json:"spec"`

	Status *UpgradePlanStatus `json:"status"`
}

func (o UpgradePlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradePlan struct{}"
	}

	return strings.Join([]string{"UpgradePlan", string(data)}, " ")
}

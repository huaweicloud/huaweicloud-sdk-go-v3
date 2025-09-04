package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserChargeTypeResult 当前租户（计费）类型
type ShowUserChargeTypeResult struct {

	// 套餐类型
	Type *string `json:"type,omitempty"`

	// 计费类型
	ChargeType *string `json:"charge_type,omitempty"`

	// 计费类型
	MainResourceList *[]ShowUserChargeTypeResultMainResourceList `json:"main_resource_list,omitempty"`
}

func (o ShowUserChargeTypeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserChargeTypeResult struct{}"
	}

	return strings.Join([]string{"ShowUserChargeTypeResult", string(data)}, " ")
}

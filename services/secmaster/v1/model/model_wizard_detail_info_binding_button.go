package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WizardDetailInfoBindingButton struct {

	// 按钮id
	ButtonId string `json:"button_id"`

	// 按钮名称
	ButtonName *string `json:"button_name,omitempty"`
}

func (o WizardDetailInfoBindingButton) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WizardDetailInfoBindingButton struct{}"
	}

	return strings.Join([]string{"WizardDetailInfoBindingButton", string(data)}, " ")
}

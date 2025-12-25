package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WizardUpdateRequestPojoBindingButton struct {

	// 按钮id
	ButtonId string `json:"button_id"`

	// 按钮名称
	ButtonName *string `json:"button_name,omitempty"`
}

func (o WizardUpdateRequestPojoBindingButton) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WizardUpdateRequestPojoBindingButton struct{}"
	}

	return strings.Join([]string{"WizardUpdateRequestPojoBindingButton", string(data)}, " ")
}

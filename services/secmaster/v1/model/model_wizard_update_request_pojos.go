package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WizardUpdateRequestPojos 更新页面请求体
type WizardUpdateRequestPojos struct {

	// 更新页面请求体列表
	LayoutWizardUpdatePojoList *[]WizardUpdateRequestPojo `json:"layout_wizard_update_pojo_list,omitempty"`
}

func (o WizardUpdateRequestPojos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WizardUpdateRequestPojos struct{}"
	}

	return strings.Join([]string{"WizardUpdateRequestPojos", string(data)}, " ")
}

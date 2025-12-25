package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupAssociationAlarmTemplateResponse Response Object
type UpdateResourceGroupAssociationAlarmTemplateResponse struct {

	// **参数解释** 资源分组ID。 **取值范围** 以\"rg\"开头，后面跟着22个字母或数字
	GroupId *string `json:"group_id,omitempty"`

	// 关联的告警模板的ID列表
	TemplateIds    *[]string `json:"template_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateResourceGroupAssociationAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupAssociationAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupAssociationAlarmTemplateResponse", string(data)}, " ")
}

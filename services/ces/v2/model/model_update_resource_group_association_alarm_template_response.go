package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupAssociationAlarmTemplateResponse Response Object
type UpdateResourceGroupAssociationAlarmTemplateResponse struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
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

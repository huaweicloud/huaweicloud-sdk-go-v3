package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupAssociationAlarmTemplateRequest Request Object
type UpdateResourceGroupAssociationAlarmTemplateRequest struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	Body *AsyncAssociateRgAndTemplatesReq `json:"body,omitempty"`
}

func (o UpdateResourceGroupAssociationAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupAssociationAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupAssociationAlarmTemplateRequest", string(data)}, " ")
}

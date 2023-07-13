package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociationResourceGroup 关联的资源分组
type AssociationResourceGroup struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	// 资源分组名称
	GroupName string `json:"group_name"`

	TemplateApplicationType *TemplateApplicationType `json:"template_application_type"`
}

func (o AssociationResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationResourceGroup struct{}"
	}

	return strings.Join([]string{"AssociationResourceGroup", string(data)}, " ")
}

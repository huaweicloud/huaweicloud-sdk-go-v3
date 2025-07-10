package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueConfigFieldsResponseBodyResultConfigFields struct {

	// 查询结果
	DisplayName *string `json:"display_name,omitempty"`

	// 是否显示
	Show *bool `json:"show,omitempty"`

	// 是否可编辑
	ShowEditable *bool `json:"show_editable,omitempty"`

	// 是否必填
	Optional *bool `json:"optional,omitempty"`

	// 是否受控
	Controlled *bool `json:"controlled,omitempty"`

	// 可编辑的角色
	IssueRoles *[]string `json:"issue_roles,omitempty"`
}

func (o IssueConfigFieldsResponseBodyResultConfigFields) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueConfigFieldsResponseBodyResultConfigFields struct{}"
	}

	return strings.Join([]string{"IssueConfigFieldsResponseBodyResultConfigFields", string(data)}, " ")
}

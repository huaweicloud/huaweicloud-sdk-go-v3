package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseRemoveInfo struct {

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 迭代Uri
	IteratorUri *string `json:"iterator_uri,omitempty"`

	// 是否移除关联的issue
	RemoveAssociateIssue *bool `json:"remove_associate_issue,omitempty"`

	// 用例列表信息
	CaseList *[]CaseIdAndTypeInfo `json:"case_list,omitempty"`
}

func (o CaseRemoveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseRemoveInfo struct{}"
	}

	return strings.Join([]string{"CaseRemoveInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalTestCaseVo struct {

	// 用例名称
	Name *string `json:"name,omitempty"`

	Owner *NameAndIdVo `json:"owner,omitempty"`

	Status *NameAndIdVo `json:"status,omitempty"`

	Result *NameAndIdVo `json:"result,omitempty"`

	Module *NameAndIdVo `json:"module,omitempty"`

	Iteration *NameAndIdVo `json:"iteration,omitempty"`

	// 用例ID
	Id *string `json:"id,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 用例描述
	Description *string `json:"description,omitempty"`

	// 用例等级
	RankId *string `json:"rank_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	ExecutionType *NameAndIdVo `json:"execution_type,omitempty"`

	TestType *IntegerIdAndNameVo `json:"test_type,omitempty"`

	CreateInfo *CreateInfoVo `json:"create_info,omitempty"`

	ExecuteInfo *ExecuteInfoVo `json:"execute_info,omitempty"`

	AssociateIssueInfo *AssociateIssueInfoVo `json:"associate_issue_info,omitempty"`

	AssociateDefectInfo *AssociateDefectInfoVo `json:"associate_defect_info,omitempty"`
}

func (o ExternalTestCaseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalTestCaseVo struct{}"
	}

	return strings.Join([]string{"ExternalTestCaseVo", string(data)}, " ")
}

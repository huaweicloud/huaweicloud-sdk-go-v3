package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试用例其他扩展信息
type ExternalServiceBizCase struct {
	// 测试用例描述信息，长度为[0-500]位字符

	Description *string `json:"description,omitempty"`
	// 执行该测试用例时需要满足的前置条件，长度为[0-500]位字符

	Preparation *string `json:"preparation,omitempty"`
	// 测试步骤，数组长度小于10

	Steps *[]ExternalServiceCaseStep `json:"steps,omitempty"`
	// 标签名称列表，数组长度小于25

	LabelList *[]string `json:"label_list,omitempty"`
	// 模块号，长度为[0-32]位字符

	ModuleId *string `json:"module_id,omitempty"`
	// 测试版本号，长度为[0-10]位字符

	TestVersionId *string `json:"test_version_id,omitempty"`
	// 迭代号，长度为[0-32]位字符

	FixVersionId *string `json:"fix_version_id,omitempty"`
	// 处理者id信息，固定长度32位字符

	AssignedId *string `json:"assigned_id,omitempty"`
	// 用例关联的需求id信息，长度为[0-32]位字符

	IssueId *string `json:"issue_id,omitempty"`
	// 测试用例状态信息，（0-新建，5-设计中，6-测试中，7-完成）

	StatusId *string `json:"status_id,omitempty"`
	// 缺陷id信息，数组长度小于50个

	DefectIdList *[]string `json:"defect_id_list,omitempty"`
}

func (o ExternalServiceBizCase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalServiceBizCase struct{}"
	}

	return strings.Join([]string{"ExternalServiceBizCase", string(data)}, " ")
}

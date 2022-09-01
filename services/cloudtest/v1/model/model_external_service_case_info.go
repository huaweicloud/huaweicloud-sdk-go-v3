package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试用例其他扩展信息
type ExternalServiceCaseInfo struct {

	// 测试用例描述信息，长度为[0-500]位字符
	Description *string `json:"description,omitempty" xml:"description"`

	// 执行该测试用例时需要满足的前置条件，长度为[0-500]位字符
	Preparation *string `json:"preparation,omitempty" xml:"preparation"`

	// 测试步骤，数组长度小于10
	Steps *[]ExternalServiceCaseStep `json:"steps,omitempty" xml:"steps"`

	// 标签名称列表，数组长度小于25
	LabelList *[]string `json:"label_list,omitempty" xml:"label_list"`

	// 模块号，长度为[0-32]位字符
	ModuleId *string `json:"module_id,omitempty" xml:"module_id"`

	// 测试版本号，长度为[0-10]位字符
	TestVersionId *string `json:"test_version_id,omitempty" xml:"test_version_id"`

	// 迭代号，长度为[0-32]位字符
	FixVersionId *string `json:"fix_version_id,omitempty" xml:"fix_version_id"`

	// 处理者id信息，固定长度32位字符
	AssignedId *string `json:"assigned_id,omitempty" xml:"assigned_id"`

	// 用例关联的需求id信息，长度为[0-32]位字符
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id"`
}

func (o ExternalServiceCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalServiceCaseInfo struct{}"
	}

	return strings.Join([]string{"ExternalServiceCaseInfo", string(data)}, " ")
}

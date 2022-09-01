package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用例其他扩展信息
type ExtendInfo struct {
	Author *ExtendAuthorInfo `json:"author,omitempty" xml:"author"`

	Updator *ExtendAuthorInfo `json:"updator,omitempty" xml:"updator"`

	Domain *AssignedUserInfo `json:"domain,omitempty" xml:"domain"`

	// 描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 前置条件
	Preparation *string `json:"preparation,omitempty" xml:"preparation"`

	// 测试步骤，数组长度小于10
	Steps *[]ExternalServiceCaseStep `json:"steps,omitempty" xml:"steps"`

	// 标签信息
	LabelList *[]AssignedUserInfo `json:"label_list,omitempty" xml:"label_list"`

	// 缺陷信息
	DefectList *[]AssignedUserInfo `json:"defect_list,omitempty" xml:"defect_list"`

	Module *AssignedUserInfo `json:"module,omitempty" xml:"module"`

	Issue *AssignedUserInfo `json:"issue,omitempty" xml:"issue"`

	// 测试版本号
	TestVersionId *string `json:"test_version_id,omitempty" xml:"test_version_id"`

	FixedVersion *AssignedUserInfo `json:"fixed_version,omitempty" xml:"fixed_version"`
}

func (o ExtendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendInfo struct{}"
	}

	return strings.Join([]string{"ExtendInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用例其他扩展信息
type ExtendInfo struct {
	Author *ExtendAuthorInfo `json:"author,omitempty"`

	Updator *ExtendAuthorInfo `json:"updator,omitempty"`

	Domain *AssignedUserInfo `json:"domain,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 前置条件
	Preparation *string `json:"preparation,omitempty"`

	// 测试步骤，数组长度小于10
	Steps *[]ExternalServiceCaseStep `json:"steps,omitempty"`

	// 标签信息
	LabelList *[]AssignedUserInfo `json:"label_list,omitempty"`

	// 缺陷信息
	DefectList *[]AssignedUserInfo `json:"defect_list,omitempty"`

	Module *AssignedUserInfo `json:"module,omitempty"`

	Issue *AssignedUserInfo `json:"issue,omitempty"`

	// 测试版本号
	TestVersionId *string `json:"test_version_id,omitempty"`

	FixedVersion *AssignedUserInfo `json:"fixed_version,omitempty"`
}

func (o ExtendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendInfo struct{}"
	}

	return strings.Join([]string{"ExtendInfo", string(data)}, " ")
}

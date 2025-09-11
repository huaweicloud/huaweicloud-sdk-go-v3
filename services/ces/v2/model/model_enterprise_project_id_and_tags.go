package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseProjectIdAndTags **参数解释** 企业项目Id或标签信息 **约束限制** 不涉及
type EnterpriseProjectIdAndTags struct {

	// **参数解释** 企业项目ID **约束限制** 不涉及 **取值范围** 由字母、数字、_和-组成，长度为[1,128]个字符 **默认取值** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tag *ResourceGroupTagRelation `json:"tag,omitempty"`
}

func (o EnterpriseProjectIdAndTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectIdAndTags struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectIdAndTags", string(data)}, " ")
}

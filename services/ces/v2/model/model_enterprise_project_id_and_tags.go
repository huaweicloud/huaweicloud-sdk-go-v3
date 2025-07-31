package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseProjectIdAndTags 企业项目Id或标签信息
type EnterpriseProjectIdAndTags struct {

	// 企业项目ID
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

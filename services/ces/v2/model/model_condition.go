package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Condition **参数解释** 匹配条件 **约束限制** 不涉及
type Condition struct {

	// **参数解释** 企业项目ID **约束限制** 不涉及 **取值范围** 由数字、字母和-组成，或者为0（默认企业项目ID） **默认取值** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	InstanceName *CombResourceName `json:"instance_name,omitempty"`

	Tag *ResourceGroupTagRelation `json:"tag,omitempty"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}

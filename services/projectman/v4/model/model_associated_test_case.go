package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociatedTestCase struct {

	// 用例ID
	CaseId *string `json:"case_id,omitempty"`

	// 用例编号
	CaseNum *string `json:"case_num,omitempty"`

	// 用例名称
	CaseName *string `json:"case_name,omitempty"`

	// 用例等级
	CaseLevel *string `json:"case_level,omitempty"`

	Status *StatusVo `json:"status,omitempty"`

	Creator *SimpleUser `json:"creator,omitempty"`

	Owner *SimpleUser `json:"owner,omitempty"`

	Project *SimpleProject `json:"project,omitempty"`

	// 是否基线
	IsBaseLine *int32 `json:"is_base_line,omitempty"`

	// 用例类型
	Type *string `json:"type,omitempty"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`
}

func (o AssociatedTestCase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedTestCase struct{}"
	}

	return strings.Join([]string{"AssociatedTestCase", string(data)}, " ")
}

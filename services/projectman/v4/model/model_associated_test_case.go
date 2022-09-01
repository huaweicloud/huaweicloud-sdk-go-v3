package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociatedTestCase struct {

	// 用例ID
	CaseId *string `json:"case_id,omitempty" xml:"case_id"`

	// 用例编号
	CaseNum *string `json:"case_num,omitempty" xml:"case_num"`

	// 用例名称
	CaseName *string `json:"case_name,omitempty" xml:"case_name"`

	// 用例等级
	CaseLevel *string `json:"case_level,omitempty" xml:"case_level"`

	Status *StatusVo `json:"status,omitempty" xml:"status"`

	Creator *SimpleUser `json:"creator,omitempty" xml:"creator"`

	Owner *SimpleUser `json:"owner,omitempty" xml:"owner"`

	Project *SimpleProject `json:"project,omitempty" xml:"project"`

	// 是否基线
	IsBaseLine *int32 `json:"is_base_line,omitempty" xml:"is_base_line"`

	// 用例类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time"`
}

func (o AssociatedTestCase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedTestCase struct{}"
	}

	return strings.Join([]string{"AssociatedTestCase", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInstanceResponseResultIssuePriority 工作项优先级
type ProcessInstanceResponseResultIssuePriority struct {

	// id
	Id *string `json:"id,omitempty"`

	// 显示名称
	DisplayValue *string `json:"display_value,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`

	// 编码
	Code *string `json:"code,omitempty"`

	// 值(拼音首字母)
	ValuePy *string `json:"value_py,omitempty"`

	// 序列
	Sequence *int32 `json:"sequence,omitempty"`

	// 层级
	Level *int32 `json:"level,omitempty"`

	// 项目ID
	DomainId *string `json:"domain_id,omitempty"`

	// 所属定义级别
	BelongDefinitionType *string `json:"belong_definition_type,omitempty"`
}

func (o ProcessInstanceResponseResultIssuePriority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultIssuePriority struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultIssuePriority", string(data)}, " ")
}

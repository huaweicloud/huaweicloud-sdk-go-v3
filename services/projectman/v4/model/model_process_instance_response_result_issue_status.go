package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInstanceResponseResultIssueStatus 工作项状态
type ProcessInstanceResponseResultIssueStatus struct {

	// 状态id
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作项的状态属性。 **取值范围**： START、IN_PROGRESS、END。
	Belonging *string `json:"belonging,omitempty"`

	// 空间id
	SpaceId *string `json:"space_id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 状态码
	Code *string `json:"code,omitempty"`

	// 定义类型
	DefinitionType *string `json:"definition_type,omitempty"`

	// 所属字段定义类型
	BelongDefinitionType *int32 `json:"belong_definition_type,omitempty"`

	// 显示值
	DisplayValue *string `json:"display_value,omitempty"`

	// 排序位置
	Position *int32 `json:"position,omitempty"`

	// 可显示
	Displayable *int32 `json:"displayable,omitempty"`

	// 可编辑
	Editable *int32 `json:"editable,omitempty"`

	// 可删除
	Deletable *int32 `json:"deletable,omitempty"`

	// 可变的
	Mutable *int32 `json:"mutable,omitempty"`

	// 状态拼音
	TitlePy *string `json:"title_py,omitempty"`

	// 状态创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 状态创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 状态最后修改时间
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 状态最后修改人
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 是否链接节点字段
	LinkageNodeFields *bool `json:"linkage_node_fields,omitempty"`
}

func (o ProcessInstanceResponseResultIssueStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultIssueStatus struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultIssueStatus", string(data)}, " ")
}

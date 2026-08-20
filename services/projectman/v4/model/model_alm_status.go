package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlmStatus struct {

	// 状态ID。
	Id *string `json:"id,omitempty"`

	// 工作项的状态属性。
	Belonging *string `json:"belonging,omitempty"`

	// 状态所属的项目空间ID。
	SpaceId *string `json:"space_id,omitempty"`

	// 状态名称。
	Name *string `json:"name,omitempty"`

	// 状态code值。
	Code *string `json:"code,omitempty"`

	// 状态定义级别，1,2,3为系统级，4为租户自定义，5为项目自定义。
	DefinitionType *string `json:"definition_type,omitempty"`

	// 状态归属定义级别，1,2,3为系统级，4为租户自定义，5为项目自定义。区别于definition_type。如果为系统级和租户自定义级，在项目中会复制一份元数据，归属于项目空间。
	BelongDefinitionType *int32 `json:"belong_definition_type,omitempty"`

	// 状态名称，和name值相同。
	DisplayValue *string `json:"display_value,omitempty"`

	// 位置顺序。
	Position *int32 `json:"position,omitempty"`

	// 是否显示。
	Displayable *int32 `json:"displayable,omitempty"`

	// 是否可编辑。
	Editable *int32 `json:"editable,omitempty"`

	// 是否可删除。
	Deletable *int32 `json:"deletable,omitempty"`

	// 是否可变，即是否为固定值。
	Mutable *int32 `json:"mutable,omitempty"`

	// 标题的拼音首字母。
	TitlePy *string `json:"title_py,omitempty"`

	// 创建人用户ID。
	CreatedBy *string `json:"created_by,omitempty"`

	// 创建时间。Unix时间戳，精度为毫秒。
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 最近修改时间。Unix时间戳，精度为毫秒。
	ModifiedDate *int64 `json:"modified_date,omitempty"`

	// 最近修改人用户ID。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 工作流配置中用于标识是否新增“节点责任人/节点结束时间”。
	LinkageNodeFields *bool `json:"linkage_node_fields,omitempty"`
}

func (o AlmStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlmStatus struct{}"
	}

	return strings.Join([]string{"AlmStatus", string(data)}, " ")
}

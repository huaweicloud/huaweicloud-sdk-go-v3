package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDpeMappingRequestBody 创建分类映射请求体
type CreateDpeMappingRequestBody struct {

	// 映射id
	Id string `json:"id"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	DataclassId string `json:"dataclass_id"`

	// 数据源
	DataSource *string `json:"data_source,omitempty"`

	// 状态(enabled，disabled)
	Status *string `json:"status,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者id
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 修改者id
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改者名称
	ModifierName *string `json:"modifier_name,omitempty"`

	Mapper *DpeMappingDetail `json:"mapper"`

	Classifier *DpeClassifyCreate `json:"classifier,omitempty"`
}

func (o CreateDpeMappingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDpeMappingRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDpeMappingRequestBody", string(data)}, " ")
}

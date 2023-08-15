package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenEntityWithExtInfoEntity 资产详情实体
type OpenEntityWithExtInfoEntity struct {

	// 属性Map<String, Object>
	Attributes *interface{} `json:"attributes,omitempty"`

	Connection *Connection `json:"connection,omitempty"`

	// 创建时间
	CreateTime float32 `json:"create_time,omitempty"`

	// 创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 资产的名称
	DisplayText *string `json:"display_text,omitempty"`

	// 资产guid
	Guid *string `json:"guid,omitempty"`

	// 相关的属性 Map<String, Object>
	RelationshipAttributes *interface{} `json:"relationship_attributes,omitempty"`

	// 资产类型
	TypeName *string `json:"type_name,omitempty"`

	// 更新人
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 更新时间
	UpdateTime float32 `json:"update_time,omitempty"`

	// 标签 List<TagHeader>
	Tags *[]TagHeader `json:"tags,omitempty"`

	// 分类名称 List<String>
	ClassificationNames *[]string `json:"classification_names,omitempty"`
}

func (o OpenEntityWithExtInfoEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenEntityWithExtInfoEntity struct{}"
	}

	return strings.Join([]string{"OpenEntityWithExtInfoEntity", string(data)}, " ")
}

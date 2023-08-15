package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenEntity 查询到的资产详情
type OpenEntity struct {

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

	// 标签
	Tags *[]TagHeader `json:"tags,omitempty"`

	// 分类 List<String> classificationNames
	ClassificationNames *[]string `json:"classification_names,omitempty"`
}

func (o OpenEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenEntity struct{}"
	}

	return strings.Join([]string{"OpenEntity", string(data)}, " ")
}

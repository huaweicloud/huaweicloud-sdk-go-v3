package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateGroupCollection struct {

	// 模板组集合id<br/>
	GroupCollectionId *string `json:"group_collection_id,omitempty" xml:"group_collection_id"`

	// 模板组集合名称<br/>
	Name *string `json:"name,omitempty" xml:"name"`

	// 模板介绍<br/>
	Description *string `json:"description,omitempty" xml:"description"`

	// 转码组列表<br/>
	TemplateGroupList *[]TemplateGroup `json:"template_group_list,omitempty" xml:"template_group_list"`
}

func (o TemplateGroupCollection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateGroupCollection struct{}"
	}

	return strings.Join([]string{"TemplateGroupCollection", string(data)}, " ")
}

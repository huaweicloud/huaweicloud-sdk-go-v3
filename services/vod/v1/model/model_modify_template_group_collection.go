package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyTemplateGroupCollection struct {

	// 模板组集合名称<br/>
	Name string `json:"name"`

	// 模板组集合ID<br/>
	CollectionId *string `json:"collection_id,omitempty"`

	// 模板组集合介绍<br/>
	Description *string `json:"description,omitempty"`

	// 模板组列表<br/>
	TemplateGroupList []string `json:"template_group_list"`
}

func (o ModifyTemplateGroupCollection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTemplateGroupCollection struct{}"
	}

	return strings.Join([]string{"ModifyTemplateGroupCollection", string(data)}, " ")
}

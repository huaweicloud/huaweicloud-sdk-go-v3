package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceGroupRequestBody struct {

	// 资源分组的名称，只能为字母、数字、汉字、-、_，最大长度为128
	GroupName string `json:"group_name"`

	// 资源分组归属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源分组创建方式，取值只能为EPS（同步企业项目）,TAG（标签动态匹配）,不传为手动添加
	Type *string `json:"type,omitempty"`

	// 标签动态匹配时的关联标签,type为TAG时必传
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`

	// 该资源分组内包含的资源来源的企业项目ID，type为EPS时必传
	AssociationEpIds *[]string `json:"association_ep_ids,omitempty"`
}

func (o CreateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequestBody", string(data)}, " ")
}

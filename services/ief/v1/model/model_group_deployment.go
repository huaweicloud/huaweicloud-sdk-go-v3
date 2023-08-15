package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupDeployment 边缘节点组绑定的应用部署
type GroupDeployment struct {

	// 应用部署uuid
	Id *string `json:"id,omitempty"`

	// 应用部署名称，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾
	Name *string `json:"name,omitempty"`

	// 应用部署描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 应用部署创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 应用部署更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o GroupDeployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupDeployment struct{}"
	}

	return strings.Join([]string{"GroupDeployment", string(data)}, " ")
}

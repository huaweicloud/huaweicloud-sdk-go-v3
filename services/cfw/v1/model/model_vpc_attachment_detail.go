package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcAttachmentDetail struct {

	// id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// vpc id
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网id
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// vpc项目id
	VpcProjectId *string `json:"vpc_project_id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o VpcAttachmentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcAttachmentDetail struct{}"
	}

	return strings.Join([]string{"VpcAttachmentDetail", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectsListInfo 项目信息
type ProjectsListInfo struct {

	// 域 ID
	DomainId *string `json:"domain_id,omitempty"`

	// 是否是域级
	IsDomain *bool `json:"is_domain,omitempty"`

	// 父项目 ID
	ParentId *string `json:"parent_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 项目ID
	Id *string `json:"id,omitempty"`

	// 是否开启
	Enabled *bool `json:"enabled,omitempty"`

	Links *SelfLinksInfo `json:"links,omitempty"`
}

func (o ProjectsListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectsListInfo struct{}"
	}

	return strings.Join([]string{"ProjectsListInfo", string(data)}, " ")
}

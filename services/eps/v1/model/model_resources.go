package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源列表
type Resources struct {
	// 企业项目ID

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// ProjectID

	ProjectId string `json:"project_id"`
	// Project名称

	ProjectName string `json:"project_name"`
	// 资源详情

	ResourceDetail *interface{} `json:"resource_detail"`
	// 资源ID

	ResourceId string `json:"resource_id"`
	// 资源名称

	ResourceName string `json:"resource_name"`
	// 资源类型

	ResourceType string `json:"resource_type"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectV3 struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 项目名

	Name string `json:"name"`
	// 区域编码

	RegionId *string `json:"region_id,omitempty"`
	// 区域名称

	RegionName *string `json:"region_name,omitempty"`
	// 管理权限

	ManagementPermission *bool `json:"management_permission,omitempty"`
	// 是否是存量项目

	IsStock *bool `json:"is_stock,omitempty"`
}

func (o ProjectV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectV3 struct{}"
	}

	return strings.Join([]string{"ProjectV3", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationBaseV3 struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name string `json:"name"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 区域id
	RegionId string `json:"region_id"`

	// 区域名称
	RegionName string `json:"region_name"`

	// 所属项目id
	ProjectId string `json:"project_id"`

	// 项目名称
	ProjectName string `json:"project_name"`

	// 应用图标
	Icon *string `json:"icon,omitempty"`
}

func (o ApplicationBaseV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationBaseV3 struct{}"
	}

	return strings.Join([]string{"ApplicationBaseV3", string(data)}, " ")
}

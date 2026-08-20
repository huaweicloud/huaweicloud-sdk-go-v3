package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModulesDetailRequest Request Object
type ListModulesDetailRequest struct {

	// 项目uuid
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 区域名
	RegionName *string `json:"region_name,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 产品线
	ProductLine *string `json:"productLine,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示数
	Limit *int32 `json:"limit,omitempty"`

	// 扩展点
	Locations []string `json:"locations"`
}

func (o ListModulesDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModulesDetailRequest struct{}"
	}

	return strings.Join([]string{"ListModulesDetailRequest", string(data)}, " ")
}

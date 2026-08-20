package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Endpoint endpoint具体信息
type Endpoint struct {
	CreatedBy *EndpointCreatorInfo `json:"created_by,omitempty"`

	// 扩展点数据
	Data *interface{} `json:"data,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 模块名称（用于搜索）
	Name *string `json:"name,omitempty"`

	// 项目uuid
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 区域名
	RegionName *string `json:"region_name,omitempty"`

	// 链接地址
	Url *string `json:"url,omitempty"`

	// 扩展点id
	Uuid *string `json:"uuid,omitempty"`
}

func (o Endpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoint struct{}"
	}

	return strings.Join([]string{"Endpoint", string(data)}, " ")
}

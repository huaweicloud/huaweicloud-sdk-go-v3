package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndPointResponse DomainsStatuses请求体
type EndPointResponse struct {
	Authorization *EndPointResponseAuthorization `json:"authorization,omitempty"`

	// uuid
	Uuid *string `json:"uuid,omitempty"`

	// 访问地址
	Url *string `json:"url,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 项目uuid
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 区域名称
	RegionName *string `json:"region_name,omitempty"`

	// 数据
	Data *interface{} `json:"data,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	CreatedBy *EndPointResponseCreatedBy `json:"created_by,omitempty"`
}

func (o EndPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndPointResponse struct{}"
	}

	return strings.Join([]string{"EndPointResponse", string(data)}, " ")
}

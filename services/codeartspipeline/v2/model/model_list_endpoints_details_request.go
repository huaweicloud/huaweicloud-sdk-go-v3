package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsDetailsRequest Request Object
type ListEndpointsDetailsRequest struct {

	// 项目uuid
	ProjectUuid string `json:"project_uuid"`

	// 区域名
	RegionName string `json:"region_name"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEndpointsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsDetailsRequest", string(data)}, " ")
}

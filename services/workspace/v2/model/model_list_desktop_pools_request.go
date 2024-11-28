package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopPoolsRequest Request Object
type ListDesktopPoolsRequest struct {

	// 桌面池名称。
	Name *string `json:"name,omitempty"`

	// 桌面池类型，DYNAMIC：动态池，STATIC：静态池。
	Type *string `json:"type,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-1000，默认值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 按照维护状态过滤
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`
}

func (o ListDesktopPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopPoolsRequest", string(data)}, " ")
}

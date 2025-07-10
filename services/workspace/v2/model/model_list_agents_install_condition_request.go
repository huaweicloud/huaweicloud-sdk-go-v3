package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentsInstallConditionRequest Request Object
type ListAgentsInstallConditionRequest struct {

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面状态。
	Status *string `json:"status,omitempty"`

	// ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 插件是否已安装。
	IsInstalled *bool `json:"is_installed,omitempty"`

	// 桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 每页显示的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAgentsInstallConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentsInstallConditionRequest struct{}"
	}

	return strings.Join([]string{"ListAgentsInstallConditionRequest", string(data)}, " ")
}

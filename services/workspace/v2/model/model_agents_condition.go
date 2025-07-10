package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentsCondition 桌面agent安装情况。
type AgentsCondition struct {

	// 桌面的desktopId。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 桌面运行状态。
	Status *string `json:"status,omitempty"`

	// 桌面的任务状态。
	TaskStatus *string `json:"task_status,omitempty"`

	// ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 桌面任务进度， 取值范围0-100以及null，null表示该桌面无任务，0-100表明该任务进度的百分比。
	Process *int32 `json:"process,omitempty"`

	// 单个桌面内的agent安装情况。
	AgentInfo *[]AgentInfo `json:"agent_info,omitempty"`
}

func (o AgentsCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentsCondition struct{}"
	}

	return strings.Join([]string{"AgentsCondition", string(data)}, " ")
}

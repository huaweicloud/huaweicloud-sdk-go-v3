package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRansomwareProtectionNodesRequest Request Object
type ListRansomwareProtectionNodesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 服务器IP地址，服务器公网IP地址
	HostIp *string `json:"host_ip,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 主机状态，包含如下3种。   - 不传参默认为全部。   - ACTIVE ：正在运行。   - SHUTOFF ：关机。
	HostStatus *string `json:"host_status,omitempty"`

	// 勒索防护状态，包含如下6种。   - closed ：未开启。   - opened ：防护中。   - opening ：开启中。   - closing ：关闭中。   - protect_failed：防护失败。   - protect_degraded：防护降级。
	RansomProtectionStatus *string `json:"ransom_protection_status,omitempty"`

	// 勒索防护策略名称
	ProtectPolicyName *string `json:"protect_policy_name,omitempty"`

	// 防护策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 防护策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// Agent状态
	AgentStatus *string `json:"agent_status,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`

	// 查询时间范围天数，1~30天，若不填，则默认查询一天
	LastDays *int64 `json:"last_days,omitempty"`
}

func (o ListRansomwareProtectionNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRansomwareProtectionNodesRequest struct{}"
	}

	return strings.Join([]string{"ListRansomwareProtectionNodesRequest", string(data)}, " ")
}

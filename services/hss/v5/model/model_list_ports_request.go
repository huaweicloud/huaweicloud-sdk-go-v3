package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsRequest Request Object
type ListPortsRequest struct {

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机IP **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 端口号 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值65535 **默认取值**: 不涉及
	Port *int32 `json:"port,omitempty"`

	// **参数解释**: 端口类型：目前包括TCP，UDP两种 **约束限制**: 不涉及 **取值范围**: - TCP: TCP类型的端口 - UDP: UDP类型的端口 **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: 端口状态 **约束限制**: 不涉及 **取值范围**: - danger: 危险端口 - unknow: 无已知危险的端口  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 主机所属的企业项目ID。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。 **约束限制**: 开通企业项目功能后才需要配置企业项目。 **取值范围**: 字符长度0-256位 **默认取值**: 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 类别，默认为host **约束限制**: 不涉及 **取值范围**: - host：主机 - container：容器 **默认取值**: host
	Category *string `json:"category,omitempty"`
}

func (o ListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsRequest struct{}"
	}

	return strings.Join([]string{"ListPortsRequest", string(data)}, " ")
}

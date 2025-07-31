package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNodesRequest Request Object
type ListContainerNodesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: Region ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: Agent状态 **约束限制**: 不涉及 **取值范围**: - online: 在线状态 - offline: 离线状态 - not_installed: 未安装Agent **默认取值**: 不涉及
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: Agent防护状态 **约束限制**: 不涉及 **取值范围**: - closed: 防护关闭状态 - opened: 防护开启状态 **默认取值**: 不涉及
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 标签: 用来识别cce集群节点和自建节点 **约束限制**: 不涉及 **取值范围**: - cce: CCE集群节点 - self: 自建集群节点 - other: 其它节点 **默认取值**: 不涉及
	ContainerTags *string `json:"container_tags,omitempty"`
}

func (o ListContainerNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNodesRequest struct{}"
	}

	return strings.Join([]string{"ListContainerNodesRequest", string(data)}, " ")
}

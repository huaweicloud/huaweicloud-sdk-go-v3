package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNodesRequest Request Object
type ListContainerNodesRequest struct {

	// Region ID
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 节点名称
	HostName *string `json:"host_name,omitempty"`

	// Agent状态，包含如下3种。   - not_installed ：未安装   - online ：在线   - offline ：离线
	AgentStatus *string `json:"agent_status,omitempty"`

	// 防护状态，包含如下2种。   - closed ：关闭   - opened ：开启
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 标签：用来识别cce容器节点和自建  - cce：cce节点  - self：自建节点  - other：其他节点
	ContainerTags *string `json:"container_tags,omitempty"`
}

func (o ListContainerNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNodesRequest struct{}"
	}

	return strings.Join([]string{"ListContainerNodesRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNodesRequest Request Object
type ListContainerNodesRequest struct {

	// 节点（服务器）名称
	HostName *string `json:"host_name,omitempty"`

	// Agent状态，包含如下3种。   - not_installed ：未安装   - online ：在线   - offline ：离线
	AgentStatus *string `json:"agent_status,omitempty"`

	// 查询返回查询容器节点列表当前页面的数，量默认10
	Limit *int32 `json:"limit,omitempty"`

	// 查询游标，初始传入0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListContainerNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNodesRequest struct{}"
	}

	return strings.Join([]string{"ListContainerNodesRequest", string(data)}, " ")
}

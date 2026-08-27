package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentInstancesResponse Response Object
type ListAgentInstancesResponse struct {

	// Agent 示例信息
	AgentInstances *[]AgentInstanceInfo `json:"agent_instances,omitempty"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAgentInstancesResponse", string(data)}, " ")
}

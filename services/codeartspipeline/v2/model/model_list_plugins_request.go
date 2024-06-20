package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginsRequest Request Object
type ListPluginsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 偏移
	Offset string `json:"offset"`

	// 大小
	Limit string `json:"limit"`

	Body *AgentPluginInfoQueryDto `json:"body,omitempty"`
}

func (o ListPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListPluginsRequest", string(data)}, " ")
}

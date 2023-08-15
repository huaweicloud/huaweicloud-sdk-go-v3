package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAgentResponse Response Object
type SearchAgentResponse struct {

	// 总页数。
	TotalPage *int32 `json:"total_page,omitempty"`

	// 总个数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 正常个数。
	OnlineCount *int32 `json:"online_count,omitempty"`

	// 心跳异常个数。
	OfflineCount *int32 `json:"offline_count,omitempty"`

	// 被关闭的个数。
	DisableCount *int32 `json:"disable_count,omitempty"`

	// agent地址列表。
	AgentInfoList  *[]InstanceInfo `json:"agent_info_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o SearchAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAgentResponse struct{}"
	}

	return strings.Join([]string{"SearchAgentResponse", string(data)}, " ")
}

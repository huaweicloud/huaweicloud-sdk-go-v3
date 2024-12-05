package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSessionsResponse Response Object
type ListInstanceSessionsResponse struct {

	// 节点的会话信息列表。
	NodeSessions   *[]ListNodeSessionsResult `json:"node_sessions,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListInstanceSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSessionsResponse", string(data)}, " ")
}

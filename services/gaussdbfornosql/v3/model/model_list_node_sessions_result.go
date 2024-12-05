package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListNodeSessionsResult struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 总会话数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 节点会话详细信息列表。
	Sessions *[]ListNodeSessionsResultSessions `json:"sessions,omitempty"`
}

func (o ListNodeSessionsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeSessionsResult struct{}"
	}

	return strings.Join([]string{"ListNodeSessionsResult", string(data)}, " ")
}

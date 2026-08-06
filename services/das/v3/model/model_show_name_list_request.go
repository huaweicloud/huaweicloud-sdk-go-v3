package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNameListRequest Request Object
type ShowNameListRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowNameListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNameListRequest struct{}"
	}

	return strings.Join([]string{"ShowNameListRequest", string(data)}, " ")
}

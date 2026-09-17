package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestSpaceRequest Request Object
type ShowLatestSpaceRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowLatestSpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestSpaceRequest struct{}"
	}

	return strings.Join([]string{"ShowLatestSpaceRequest", string(data)}, " ")
}

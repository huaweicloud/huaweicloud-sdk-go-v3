package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNodesInfoRequest Request Object
type ShowInstanceNodesInfoRequest struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	AllNodes *string `json:"all_nodes,omitempty"`

	ShowHiddenNodes *string `json:"show_hidden_nodes,omitempty"`
}

func (o ShowInstanceNodesInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNodesInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceNodesInfoRequest", string(data)}, " ")
}

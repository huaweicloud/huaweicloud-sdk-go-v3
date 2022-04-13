package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNodeRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`
	// DDM节点ID

	NodeId string `json:"node_id"`
}

func (o ShowNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowNodeRequest", string(data)}, " ")
}

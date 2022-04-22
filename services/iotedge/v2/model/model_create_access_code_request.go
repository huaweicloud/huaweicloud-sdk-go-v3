package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAccessCodeRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 设备ID
	DeviceId string `json:"device_id"`
}

func (o CreateAccessCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessCodeRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessCodeRequest", string(data)}, " ")
}

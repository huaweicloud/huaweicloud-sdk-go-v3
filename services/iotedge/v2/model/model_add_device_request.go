package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDeviceRequest Request Object
type AddDeviceRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	Body *AddDeviceRequestBody `json:"body,omitempty"`
}

func (o AddDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeviceRequest struct{}"
	}

	return strings.Join([]string{"AddDeviceRequest", string(data)}, " ")
}
